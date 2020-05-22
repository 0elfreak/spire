package bundle

import (
	"context"

	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/spire/pkg/server/api/rpccontext"
	"github.com/spiffe/spire/pkg/server/plugin/datastore"
	"github.com/spiffe/spire/proto/spire-next/api/server/bundle/v1"
	"github.com/spiffe/spire/proto/spire-next/types"
	"github.com/spiffe/spire/proto/spire/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterService registers the bundle service on the gRPC server.
func RegisterService(s *grpc.Server, c *Config) {
	srv := service{c: c}
	bundle.RegisterBundleServer(s, srv)
}

type service struct {
	c *Config
}

type Config struct {
	ds datastore.DataStore
	td spiffeid.TrustDomain
}

func (s service) GetBundle(ctx context.Context, req *bundle.GetBundleRequest) (*types.Bundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBundle not implemented")
}

func (s service) AppendBundle(ctx context.Context, req *bundle.AppendBundleRequest) (*types.Bundle, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendBundle not implemented")
}

func (s service) ListFederatedBundles(ctx context.Context, req *bundle.ListFederatedBundlesRequest) (*bundle.ListFederatedBundlesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFederatedBundles not implemented")
}

func (s service) GetFederatedBundle(ctx context.Context, req *bundle.GetFederatedBundleRequest) (*types.Bundle, error) {
	log := rpccontext.Logger(ctx).WithField("RPC", "GetFedertedBundle")

	if !(rpccontext.CallerIsLocal(ctx) || rpccontext.CallerIsAdmin(ctx) || rpccontext.CallerIsAgent(ctx)) {
		log.Errorf("Permission denied: the caller must be local or present an admin or an active agent X509-SVID")
		return nil, status.Errorf(codes.PermissionDenied, "the caller must be local or present an admin or an active agent X509-SVID")
	}

	td, err := spiffeid.TrustDomainFromString(req.TrustDomain)
	if err != nil {
		log.Errorf("Trust domain argument is not a valid SPIFFE ID: %q", req.TrustDomain)
		return nil, status.Errorf(codes.InvalidArgument, "trust domain argument is not a valid SPIFFE ID: %q", req.TrustDomain)
	}

	if s.c.td.Compare(td) != 0 {
		log.Errorf("%q is this server own trust domain, use GetBundle RPC instead", td.String())
		return nil, status.Errorf(codes.InvalidArgument, "%q is this server own trust domain, use GetBundle RPC instead", td.String())
	}

	dsResp, err := s.c.ds.FetchBundle(ctx, &datastore.FetchBundleRequest{
		TrustDomainId: td.IDString(),
	})
	if err != nil {
		log.Errorf("Failed to fetch bundle: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to fetch bundle: %v", err)
	}

	if dsResp.Bundle == nil {
		log.Errorf("Bundle for %q not found", req.TrustDomain)
		return nil, status.Errorf(codes.NotFound, "bundle for %q not found", req.TrustDomain)
	}

	return applyMask(dsResp.Bundle, req.OutputMask), nil
}

func (s service) BatchCreateFederatedBundle(ctx context.Context, req *bundle.BatchCreateFederatedBundleRequest) (*bundle.BatchCreateFederatedBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchCreateFederatedBundle not implemented")
}

func (s service) BatchUpdateFederatedBundle(ctx context.Context, req *bundle.BatchUpdateFederatedBundleRequest) (*bundle.BatchUpdateFederatedBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchUpdateFederatedBundle not implemented")
}

func (s service) BatchSetFederatedBundle(ctx context.Context, req *bundle.BatchSetFederatedBundleRequest) (*bundle.BatchSetFederatedBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchSetFederatedBundle not implemented")
}

func (s service) BatchDeleteFederatedBundle(ctx context.Context, req *bundle.BatchDeleteFederatedBundleRequest) (*bundle.BatchDeleteFederatedBundleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDeleteFederatedBundle not implemented")
}

func applyMask(b *common.Bundle, mask *types.BundleMask) *types.Bundle {
	out := &types.Bundle{}
	if mask.TrustDomainId {
		out.TrustDomainId.TrustDomain = b.TrustDomainId
	}

	if mask.RefreshHint {
		out.RefreshHint = b.RefreshHint
	}

	if mask.SequenceNumber {
		//TODO: filter sequence numbers when SPIRE supports them
		out.SequenceNumber = 0
	}

	if mask.X509Authorities {
		authorities := []*types.X509Certificate{}
		for _, rootCA := range b.RootCas {
			authorities = append(authorities, &types.X509Certificate{
				Asn1: rootCA.DerBytes,
			})
		}
		out.X509Authorities = authorities
	}

	if mask.JwtAuthorities {
		authorities := []*types.JWTKey{}
		for _, JWTSigningKey := range b.JwtSigningKeys {
			authorities = append(authorities, &types.JWTKey{
				PublicKey: JWTSigningKey.PkixBytes,
				KeyId:     JWTSigningKey.Kid,
				ExpiresAt: JWTSigningKey.NotAfter,
			})
		}
		out.JwtAuthorities = authorities
	}

	return out
}
