package uaa

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"

	jwtv3 "github.com/dgrijalva/jwt-go"
	"github.com/lestrrat-go/jwx/jwk"
)

type KeyFetcher struct {
	UAAURL  string
	Fetcher Fetcher
}

type Fetcher interface {
	Fetch(ctx context.Context, url string) (jwk.Set, error)
}

func NewKeyFetcher(ctx context.Context, uaaBaseURL string) *KeyFetcher {
	uaaURL := fmt.Sprintf("%s/token_keys", uaaBaseURL)
	autoRefresh := jwk.NewAutoRefresh(ctx)
	autoRefresh.Configure(uaaURL)
	return &KeyFetcher{
		UAAURL:  uaaURL,
		Fetcher: autoRefresh,
	}
}

func (ukf *KeyFetcher) Fetch(token *jwtv3.Token) (interface{}, error) {
	keySet, err := ukf.Fetcher.Fetch(context.Background(), ukf.UAAURL)
	if err != nil {
		return nil, err
	}
	keyID, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("expecting JWT header to have a key ID in the kid field")
	}
	key, found := keySet.LookupKeyID(keyID)
	if !found {
		return nil, fmt.Errorf("unable to find key %q", keyID)
	}
	var pubkey rsa.PublicKey
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("unable to decode public key: %w", err)
	}

	return pubkey, nil
}
