package jwtsvid

import (
	"context"
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/suite"
)

const (
	fakeSpiffeID = "spiffe://example.org/blog"
)

var (
	ctx           = context.Background()
	fakeAudience  = []string{"AUDIENCE"}
	fakeAudiences = []string{"AUDIENCE1", "AUDIENCE2"}

	keyPEM = []byte(`-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgt/OIyb8Ossz/5bNk
XtnzFe1T2d0D9quX9Loi1O55b8yhRANCAATDe/2d6z+P095I3dIkocKr4b3zAy+1
qQDuoXqa8i3YOPk5fLib4ORzqD9NJFcrKjI+LLtipQe9yu/eY1K0yhBa
-----END PRIVATE KEY-----
`)
)

func TestToken(t *testing.T) {
	suite.Run(t, new(TokenSuite))
}

type TokenSuite struct {
	suite.Suite

	key    *ecdsa.PrivateKey
	bundle TrustBundle
}

func (s *TokenSuite) SetupTest() {
	s.key = s.loadKey(keyPEM)
	s.bundle = NewTrustBundle("example.org", map[string]crypto.PublicKey{
		"kid": s.key.Public(),
	})
}

func (s *TokenSuite) loadKey(pemBytes []byte) *ecdsa.PrivateKey {
	block, rest := pem.Decode(pemBytes)
	s.Require().Empty(rest)
	s.Require().NotNil(block)
	rawKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	s.Require().NoError(err)
	key, ok := rawKey.(*ecdsa.PrivateKey)
	s.Require().True(ok)
	return key
}

func (s *TokenSuite) TestSignAndValidate() {
	token, err := SignToken(fakeSpiffeID, fakeAudience, time.Now().Add(time.Hour), s.key, "kid")
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	claims, err := ValidateToken(ctx, token, s.bundle, fakeAudience[0])
	s.Require().NoError(err)
	s.Require().NotEmpty(claims)
}

func (s *TokenSuite) TestSignAndValidateWithAudienceList() {
	token, err := SignToken(fakeSpiffeID, fakeAudiences, time.Now().Add(time.Hour), s.key, "kid")
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	claims, err := ValidateToken(ctx, token, s.bundle, fakeAudiences[0])
	s.Require().NoError(err)
	s.Require().NotEmpty(claims)
}

func (s *TokenSuite) TestSignWithNoExpiration() {
	_, err := SignToken(fakeSpiffeID, fakeAudience, time.Time{}, s.key, "kid")
	s.Require().EqualError(err, "expiration is required")
}

func (s *TokenSuite) TestSignInvalidSpiffeID() {
	// missing ID
	_, err := SignToken("", fakeAudience, time.Now(), s.key, "kid")
	s.requireErrorContains(err, "is not a valid workload SPIFFE ID: SPIFFE ID is empty")

	// not a spiffe ID
	_, err = SignToken("sparfe://example.org", fakeAudience, time.Now(), s.key, "kid")
	s.requireErrorContains(err, "is not a valid workload SPIFFE ID: invalid scheme")
}

func (s *TokenSuite) TestSignNoAudience() {
	_, err := SignToken(fakeSpiffeID, nil, time.Now().Add(time.Hour), s.key, "kid")
	s.Require().EqualError(err, "audience is required")
}

func (s *TokenSuite) TestValidateBadAlgorithm() {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, err := token.SignedString([]byte("BLAH"))
	s.Require().NoError(err)

	claims, err := ValidateToken(ctx, tokenString, s.bundle, fakeAudience[0])
	s.Require().EqualError(err, "unexpected token signature algorithm: HS256")
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateMissingThumbprint() {
	token := jwt.New(jwt.SigningMethodES256)
	tokenString, err := token.SignedString(s.key)
	s.Require().NoError(err)

	claims, err := ValidateToken(ctx, tokenString, s.bundle, fakeAudience[0])
	s.Require().EqualError(err, "token missing key id")
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateExpiredToken() {
	token, err := SignToken(fakeSpiffeID, fakeAudience, time.Now().Add(-time.Hour), s.key, "kid")
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	claims, err := ValidateToken(ctx, token, s.bundle, fakeAudience[0])
	s.Require().EqualError(err, "Token is expired")
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateNoSubject() {
	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = "kid"
	tokenString, err := token.SignedString(s.key)
	s.Require().NoError(err)

	claims, err := ValidateToken(ctx, tokenString, s.bundle, "FOO")
	s.Require().EqualError(err, "token missing subject claim")
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateSubjectNotForDomain() {
	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = "kid"
	token.Claims = jwt.MapClaims{
		"sub": "spiffe://other.org",
	}
	tokenString, err := token.SignedString(s.key)
	s.Require().NoError(err)

	claims, err := ValidateToken(ctx, tokenString, s.bundle, "FOO")
	s.Require().EqualError(err, `token has in invalid subject claim: "spiffe://other.org" does not belong to trust domain "example.org"`)
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateNoAudience() {
	token := jwt.New(jwt.SigningMethodES256)
	token.Header["kid"] = "kid"
	token.Claims = jwt.MapClaims{
		"sub": "spiffe://example.org/blog",
	}
	tokenString, err := token.SignedString(s.key)
	s.Require().NoError(err)

	claims, err := ValidateToken(ctx, tokenString, s.bundle, "FOO")
	s.Require().EqualError(err, "token missing audience claim")
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateUnexpectedAudience() {
	token, err := SignToken(fakeSpiffeID, fakeAudience, time.Now().Add(time.Hour), s.key, "kid")
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	claims, err := ValidateToken(ctx, token, s.bundle, "FOO")
	s.Require().EqualError(err, `expected audience "FOO" (audience="AUDIENCE")`)
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateUnexpectedAudienceList() {
	token, err := SignToken(fakeSpiffeID, fakeAudiences, time.Now().Add(time.Hour), s.key, "kid")
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	claims, err := ValidateToken(ctx, token, s.bundle, "AUDIENCE3")
	s.Require().EqualError(err, `expected audience "AUDIENCE3" (audience=["AUDIENCE1" "AUDIENCE2"])`)
	s.Require().Nil(claims)
}

func (s *TokenSuite) TestValidateCertificateNotFound() {
	token, err := SignToken(fakeSpiffeID, fakeAudience, time.Now().Add(time.Hour), s.key, "whatever")
	s.Require().NoError(err)
	s.Require().NotEmpty(token)

	claims, err := ValidateToken(ctx, token, s.bundle, fakeAudience[0])
	s.Require().EqualError(err, "public key not found in trust bundle")
	s.Require().Nil(claims)
}

func (s *TokenSuite) requireErrorContains(err error, contains string) {
	s.Require().Error(err)
	s.Require().Contains(err.Error(), contains)
}
