package svid

import (
	"context"
	"crypto/x509"
	"time"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/common/x509util"
	"github.com/spiffe/spire/pkg/server/api"
	"github.com/spiffe/spire/pkg/server/api/rpccontext"
	"github.com/spiffe/spire/pkg/server/ca"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Config is the service configuration
type Config struct {
	ServerCA    ca.ServerCA
	TrustDomain spiffeid.TrustDomain
}

// Service is the SVID service interface
type Service interface {
	MintX509SVID(ctx context.Context, csr *x509.CertificateRequest, ttl time.Duration) (*api.X509SVID, error)
	MintJWTSVID(ctx context.Context, id spiffeid.ID, audience []string, ttl time.Duration) (*api.JWTSVID, error)
	NewX509SVID(ctx context.Context, entryID string, csr *x509.CertificateRequest) (*api.X509SVID, error)
	NewJWTSVID(ctx context.Context, entryID string, audience []string) (*api.JWTSVID, error)
}

// New creates a new SVID service
func New(config Config) Service {
	return &service{
		ServerCA:    config.ServerCA,
		TrustDomain: config.TrustDomain,
	}
}

type service struct {
	ServerCA    ca.ServerCA
	TrustDomain spiffeid.TrustDomain
}

func (s *service) MintX509SVID(ctx context.Context, csr *x509.CertificateRequest, ttl time.Duration) (*api.X509SVID, error) {
	log := rpccontext.Logger(ctx)

	if err := csr.CheckSignature(); err != nil {
		log.WithError(err).Error("Invalid CSR: signature verify failed")
		return nil, status.Errorf(codes.InvalidArgument, "invalid CSR: signature verify failed")
	}

	if len(csr.URIs) != 1 {
		log.Error("Invalid CSR: a valid URI is required")
		return nil, status.Errorf(codes.InvalidArgument, "invalid CSR: a valid URI is required")
	}

	spiffeID, err := spiffeid.FromURI(csr.URIs[0])
	if err != nil {
		log.WithError(err).Error("Invalid CSR: a valid SPIFFE ID is expected")
		return nil, status.Errorf(codes.InvalidArgument, "invalid CSR: a valid SPIFFE ID is expected: %v", err)
	}

	if !spiffeID.MemberOf(s.TrustDomain) {
		log.Error("Invalid CSR: SPIFFE ID is not member of the server trust domain")
		return nil, status.Error(codes.InvalidArgument, "invalid CSR: SPIFFE ID is not member of the server trust domain")
	}

	for _, dnsName := range csr.DNSNames {
		if err := x509util.ValidateDNS(dnsName); err != nil {
			log.WithError(err).Error("Invalid CSR: DNS name is not valid")
			return nil, status.Errorf(codes.InvalidArgument, "invalid CSR: DNS name is not valid: %v", err)
		}
	}

	svid, err := s.ServerCA.SignX509SVID(ctx, ca.X509SVIDParams{
		SpiffeID:  spiffeID.String(),
		PublicKey: csr.PublicKey,
		TTL:       ttl,
		DNSList:   csr.DNSNames,
		Subject:   csr.Subject,
	})
	if err != nil {
		log.WithError(err).Error("Failed to sign X509-SVID")
		return nil, status.Errorf(codes.Internal, "failed to sign X509-SVID: %v", err)
	}

	return &api.X509SVID{
		ID:        spiffeID,
		CertChain: svid,
		ExpiresAt: svid[0].NotAfter.UTC(),
	}, nil
}

func (s *service) MintJWTSVID(ctx context.Context, id spiffeid.ID, audience []string, ttl time.Duration) (*api.JWTSVID, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *service) NewX509SVID(ctx context.Context, entryID string, csr *x509.CertificateRequest) (*api.X509SVID, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}

func (s *service) NewJWTSVID(ctx context.Context, entryID string, audience []string) (*api.JWTSVID, error) {
	return nil, status.Error(codes.Unimplemented, "not implemented")
}
