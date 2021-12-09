package info

import (
	"net/url"

	"github.com/cloudfoundry/go-cf-api/internal/config"
)

type Link struct {
	HREF string            `json:"href"`
	Meta map[string]string `json:"meta,omitempty"`
}

func ExternalURL(endpoint string, conf *config.CfAPIConfig) string {
	externalURL := url.URL{
		Scheme: conf.ExternalProtocol,
		Host:   conf.ExternalDomain,
		Path:   endpoint,
	}
	return externalURL.String()
}
