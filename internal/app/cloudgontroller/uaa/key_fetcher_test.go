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
)

func TestNewUaaKeyFetcher(t *testing.T) {
	uaaBaseURL := "http://uaa.example.com"
	ukf := NewUaaKeyFetcher(context.TODO(), uaaBaseURL)

	assert.Equal(t, "http://uaa.example.com/token_keys", ukf.uaaURL)
	assert.IsType(t, &jwk.AutoRefresh{}, ukf.fetcher)
}

type MockFetcher struct {
	mock.Mock
}

func (m *MockFetcher) Fetch(ctx context.Context, url string) (jwk.Set, error) {
	args := m.Called(ctx, url)
	return args.Get(0).(jwk.Set), args.Error(1)
}

func TestUaaKeyFetcherFetchWithSingleKey(t *testing.T) {
	fetcher := &MockFetcher{}
	set := jwk.NewSet()
	rsaKey := rsa.PublicKey{N: big.NewInt(rand.Int63()), E: rand.Int()}
	publicKey := jwk.NewRSAPublicKey()
	publicKey.FromRaw(&rsaKey)
	publicKey.Set(jwk.KeyIDKey, "key-id")

	set.Add(publicKey)
	ukf := UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: fetcher,
	}
	fetcher.On("Fetch", mock.Anything, mock.Anything).Return(set, nil)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": publicKey.KeyID()}}
	key, err := ukf.Fetch(&token)
	assert.NoError(t, err)
	assert.EqualValues(t, rsaKey, key)
}

func TestUaaKeyFetcherFetchWithMultipleKeys(t *testing.T) {
	fetcher := &MockFetcher{}
	set := jwk.NewSet()
	rsaKey1 := rsa.PublicKey{N: big.NewInt(rand.Int63()), E: rand.Int()}
	rsaKey2 := rsa.PublicKey{N: big.NewInt(rand.Int63()), E: rand.Int()}
	publicKey1, publicKey2 := jwk.NewRSAPublicKey(), jwk.NewRSAPublicKey()
	publicKey1.FromRaw(&rsaKey1)
	publicKey1.Set(jwk.KeyIDKey, "key-id-1")
	publicKey2.FromRaw(&rsaKey2)
	publicKey2.Set(jwk.KeyIDKey, "key-id-2")
	set.Add(publicKey1)
	set.Add(publicKey2)
	ukf := UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: fetcher,
	}
	fetcher.On("Fetch", mock.Anything, mock.Anything).Return(set, nil)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": publicKey2.KeyID()}}
	key, err := ukf.Fetch(&token)
	assert.NoError(t, err)
	assert.EqualValues(t, rsaKey2, key)
}

func TestUaaKeyFetcherFetchNoMatchingKey(t *testing.T) {
	fetcher := &MockFetcher{}
	set := jwk.NewSet()
	ukf := UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: fetcher,
	}
	fetcher.On("Fetch", mock.Anything, mock.Anything).Return(set, nil)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": "abcd"}}
	_, err := ukf.Fetch(&token)
	assert.EqualError(t, err, `unable to find key "abcd"`)
}

func TestUaaKeyFetcherFetchNoKeyIDHeader(t *testing.T) {
	fetcher := &MockFetcher{}
	set := jwk.NewSet()
	ukf := UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: fetcher,
	}
	fetcher.On("Fetch", mock.Anything, mock.Anything).Return(set, nil)

	token := jwtv3.Token{}
	_, err := ukf.Fetch(&token)
	assert.EqualError(t, err, `expecting JWT header to have a key ID in the kid field`)
}

func TestUaaKeyFetcherFetchReturnsError(t *testing.T) {
	fetcher := &MockFetcher{}
	ukf := UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: fetcher,
	}
	fetcher.On("Fetch", mock.Anything, mock.Anything).Return(jwk.NewSet(), errors.New("couldn't fetch keys from UAA"))

	token := jwtv3.Token{}
	_, err := ukf.Fetch(&token)
	assert.EqualError(t, err, "couldn't fetch keys from UAA")
}

func TestUaaKeyFetcherFetchKeyIsNotRSA(t *testing.T) {
	fetcher := &MockFetcher{}
	set := jwk.NewSet()
	ecdsaKey := ecdsa.PublicKey{X: big.NewInt(rand.Int63()), Y: big.NewInt(rand.Int63()), Curve: elliptic.P256()}
	publicKey := jwk.NewECDSAPublicKey()
	publicKey.FromRaw(&ecdsaKey)
	publicKey.Set(jwk.KeyIDKey, "key-id")

	set.Add(publicKey)
	ukf := UaaKeyFetcher{
		uaaURL:  "http://uaa.example.com/token_keys",
		fetcher: fetcher,
	}
	fetcher.On("Fetch", mock.Anything, mock.Anything).Return(set, nil)

	token := jwtv3.Token{Header: map[string]interface{}{"kid": publicKey.KeyID()}}
	_, err := ukf.Fetch(&token)
	assert.EqualError(t, err, "unable to decode public key: argument to AssignIfCompatible() must be compatible with *ecdsa.PublicKey (was *rsa.PublicKey)")
}
