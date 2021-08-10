// +build unit

package uaa

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"errors"
	"math/big"
	"math/rand"
	"testing"

	jwtv3 "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestNewUaaKeyFetcher(t *testing.T) {
	uaaBaseURL := "http://uaa.example.com"
	ukf := NewUaaKeyFetcher(context.TODO(), uaaBaseURL)

	assert.Equal(t, "http://uaa.example.com/token_keys", ukf.uaaURL)
	assert.IsType(t, &jwk.AutoRefresh{}, ukf.fetcher)
}

func (suite *UaaKeyFetcherFetchSuite) TestUaaKeyFetcherFetchWithSingleKey() {
	rsaKey, publicKey := suite.generateRSAKey("key-id")
	suite.KeySet.Add(publicKey)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": publicKey.KeyID()}}
	key, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.NoError(err)
	suite.EqualValues(rsaKey, key)
}

func (suite *UaaKeyFetcherFetchSuite) TestUaaKeyFetcherFetchWithMultipleKeys() {
	_, publicKey1 := suite.generateRSAKey("key-id-1")
	rsaKey2, publicKey2 := suite.generateRSAKey("key-id-2")

	suite.KeySet.Add(publicKey1)
	suite.KeySet.Add(publicKey2)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": publicKey2.KeyID()}}
	key, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.NoError(err)
	suite.EqualValues(rsaKey2, key)
}

func (suite *UaaKeyFetcherFetchSuite) TestUaaKeyFetcherFetchNoMatchingKey() {
	token := jwtv3.Token{Header: map[string]interface{}{"kid": "abcd"}}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, `unable to find key "abcd"`)
}

func (suite *UaaKeyFetcherFetchSuite) TestUaaKeyFetcherFetchNoKeyIDHeader() {
	token := jwtv3.Token{}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, `expecting JWT header to have a key ID in the kid field`)
}

func (suite *UaaKeyFetcherFetchSuite) TestUaaKeyFetcherFetchReturnsError() {
	suite.MockFetcher.ExpectedCalls = nil
	suite.MockFetcher.On("Fetch", mock.Anything, mock.Anything).Return(jwk.NewSet(), errors.New("couldn't fetch keys from UAA"))

	token := jwtv3.Token{}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, "couldn't fetch keys from UAA")
}

func (suite *UaaKeyFetcherFetchSuite) TestUaaKeyFetcherFetchKeyIsNotRSA() {
	ecdsaKey := ecdsa.PublicKey{X: big.NewInt(rand.Int63()), Y: big.NewInt(rand.Int63()), Curve: elliptic.P256()}
	publicKey := jwk.NewECDSAPublicKey()
	publicKey.FromRaw(&ecdsaKey)
	publicKey.Set(jwk.KeyIDKey, "key-id")

	suite.KeySet.Add(publicKey)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": publicKey.KeyID()}}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, "unable to decode public key: argument to AssignIfCompatible() must be compatible with *ecdsa.PublicKey (was *rsa.PublicKey)")
}

func (suite *UaaKeyFetcherFetchSuite) SetupTest() {
	suite.MockFetcher = MockFetcher{}
	suite.KeySet = jwk.NewSet()
	suite.UaaKeyFetcher = UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: &suite.MockFetcher,
	}
	suite.MockFetcher.On("Fetch", mock.Anything, mock.Anything).Return(suite.KeySet, nil)
}

func (suite *UaaKeyFetcherFetchSuite) After(suiteName, testName string) {
	suite.KeySet.Clear()
}

type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) Fetch(ctx context.Context, url string) (jwk.Set, error) {
	args := m.Called(ctx, url)
	return args.Get(0).(jwk.Set), args.Error(1)
}

type UaaKeyFetcherFetchSuite struct {
	suite.Suite
	MockFetcher   MockFetcher
	KeySet        jwk.Set
	UaaKeyFetcher UaaKeyFetcher
}

func TestUaaKeyFetcherFetchSuite(t *testing.T) {
	suite.Run(t, new(UaaKeyFetcherFetchSuite))
}

func (suite *UaaKeyFetcherFetchSuite) generateRSAKey(id string) (rsa.PublicKey, jwk.RSAPublicKey) {
	rsaKey := rsa.PublicKey{N: big.NewInt(rand.Int63()), E: rand.Int()}
	publicKey := jwk.NewRSAPublicKey()
	publicKey.FromRaw(&rsaKey)
	publicKey.Set(jwk.KeyIDKey, id)
	return rsaKey, publicKey
}
