// +build unit

//nolint:gosec // These are just tests so insecure random number generation is fine
package uaa_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rsa"
	"errors"
	"math/big"
	"math/rand"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
	. "github.tools.sap/cloudfoundry/cloudgontroller/internal/uaa"
)

func TestNewKeyFetcher(t *testing.T) {
	t.Parallel()
	uaaBaseURL := "http://uaa.example.com"
	ukf, err := NewKeyFetcher(context.TODO(), config.UaaConfig{URL: uaaBaseURL})

	assert.NoError(t, err)
	assert.Equal(t, "http://uaa.example.com/token_keys", ukf.UAAURL)
	assert.IsType(t, &jwk.AutoRefresh{}, ukf.Fetcher)
}

func (suite *KeyFetcherSuite) TestUaaKeyFetcherFetchWithSingleKey() {
	rsaKey, publicKey := suite.generateRSAKey("key-id")
	suite.KeySet.Add(publicKey)

	token := jwt.Token{Header: map[string]interface{}{"kid": publicKey.KeyID()}}
	key, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.NoError(err)
	suite.EqualValues(rsaKey, key)
}

func (suite *KeyFetcherSuite) TestUaaKeyFetcherFetchWithMultipleKeys() {
	_, publicKey1 := suite.generateRSAKey("key-id-1")
	rsaKey2, publicKey2 := suite.generateRSAKey("key-id-2")

	suite.KeySet.Add(publicKey1)
	suite.KeySet.Add(publicKey2)

	token := jwt.Token{Header: map[string]interface{}{"kid": publicKey2.KeyID()}}
	key, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.NoError(err)
	suite.EqualValues(rsaKey2, key)
}

func (suite *KeyFetcherSuite) TestUaaKeyFetcherFetchNoMatchingKey() {
	token := jwt.Token{Header: map[string]interface{}{"kid": "abcd"}}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, `unable to find key "abcd"`)
}

func (suite *KeyFetcherSuite) TestUaaKeyFetcherFetchNoKeyIDHeader() {
	token := jwt.Token{}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, `expecting JWT header to have a key ID in the kid field`)
}

func (suite *KeyFetcherSuite) TestUaaKeyFetcherFetchReturnsError() {
	suite.MockFetcher.ExpectedCalls = nil
	suite.MockFetcher.On("Fetch", mock.Anything, mock.Anything).Return(jwk.NewSet(), errors.New("could not fetch keys from UAA"))

	token := jwt.Token{}
	_, err := suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, "could not fetch keys from UAA")
}

func (suite *KeyFetcherSuite) TestUaaKeyFetcherFetchKeyIsNotRSA() {
	ecdsaKey := ecdsa.PublicKey{X: big.NewInt(rand.Int63()), Y: big.NewInt(rand.Int63()), Curve: elliptic.P256()}
	publicKey := jwk.NewECDSAPublicKey()
	err := publicKey.FromRaw(&ecdsaKey)
	suite.NoError(err)
	err = publicKey.Set(jwk.KeyIDKey, "key-id")
	suite.NoError(err)

	suite.KeySet.Add(publicKey)

	token := jwt.Token{Header: map[string]interface{}{"kid": publicKey.KeyID()}}
	_, err = suite.UaaKeyFetcher.Fetch(&token)
	suite.EqualError(err, "could not use key from UAA - keys from UAA must be RSA")
}

func (suite *KeyFetcherSuite) SetupTest() {
	suite.MockFetcher = MockFetcher{}
	suite.KeySet = jwk.NewSet()
	suite.UaaKeyFetcher = KeyFetcher{
		UAAURL:  "http://uaa.example.com/token_keys",
		Fetcher: &suite.MockFetcher,
	}
	suite.MockFetcher.On("Fetch", mock.Anything, mock.Anything).Return(suite.KeySet, nil)
}

func (suite *KeyFetcherSuite) After(suiteName, testName string) {
	suite.KeySet.Clear()
}

type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) Fetch(ctx context.Context, url string) (jwk.Set, error) {
	args := m.Called(ctx, url)
	return args.Get(0).(jwk.Set), args.Error(1)
}

type KeyFetcherSuite struct {
	suite.Suite
	MockFetcher   MockFetcher
	KeySet        jwk.Set
	UaaKeyFetcher KeyFetcher
}

func TestUaaKeyFetcherFetchSuite(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(KeyFetcherSuite))
}

func (suite *KeyFetcherSuite) generateRSAKey(id string) (*rsa.PublicKey, jwk.RSAPublicKey) {
	rsaKey := &rsa.PublicKey{N: big.NewInt(rand.Int63()), E: rand.Int()}
	publicKey := jwk.NewRSAPublicKey()
	err := publicKey.FromRaw(rsaKey)
	suite.NoError(err)
	err = publicKey.Set(jwk.KeyIDKey, id)
	suite.NoError(err)
	return rsaKey, publicKey
}
