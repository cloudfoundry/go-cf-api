package uaa

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/lestrrat-go/jwx/jwk"
	promConfig "github.com/prometheus/common/config"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

type KeyFetcher struct {
	UAAURL  string
	Fetcher Fetcher
}

type Fetcher interface {
	Fetch(ctx context.Context, url string) (jwk.Set, error)
}

func NewKeyFetcher(ctx context.Context, uaaConf config.UaaConfig) (*KeyFetcher, error) {
	uaaURL := fmt.Sprintf("%s/token_keys", uaaConf.URL)
	autoRefresh := jwk.NewAutoRefresh(ctx)
	client, err := promConfig.NewClientFromConfig(uaaConf.Client, "uaa")
	if err != nil {
		return nil, err
	}
	autoRefresh.Configure(uaaURL, jwk.WithHTTPClient(client))
	return &KeyFetcher{
		UAAURL:  uaaURL,
		Fetcher: autoRefresh,
	}, nil
}

func (ukf *KeyFetcher) Fetch(token *jwt.Token) (interface{}, error) {
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
	var pubkey interface{}
	if err := key.Raw(&pubkey); err != nil {
		return nil, fmt.Errorf("unable to decode public key: %w", err)
	}

	if _, ok := pubkey.(*rsa.PublicKey); !ok {
		return nil, fmt.Errorf("could not use key from UAA - keys from UAA must be RSA")
	}

	return pubkey, nil
}
