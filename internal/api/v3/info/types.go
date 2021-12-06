package info

import (
	"github.tools.sap/cloudfoundry/cloudgontroller/internal/config"
	"net/url"
)

type Link struct {
	HREF string            `json:"href"`
	Meta map[string]string `json:"meta,omitempty"`
}

func ExternalURL(endpoint string, config *config.CfApiConfig) string {
	externalUrl := url.URL{
		Scheme: config.ExternalProtocol,
		Host:   config.ExternalDomain,
		Path:   endpoint,
	}
	return externalUrl.String()
}
