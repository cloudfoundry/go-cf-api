package info

import (
	"net/url"

	"github.tools.sap/cloudfoundry/cloudgontroller/internal/app/cloudgontroller/config"
)

type Controller struct {
	Info             config.InfoConfig
	URLs             config.URLs
	AppSSH           config.AppSSH
	ExternalDomain   string
	ExternalProtocol string
}

type Link struct {
	HREF string            `json:"href"`
	Meta map[string]string `json:"meta,omitempty"`
}

func (cont *Controller) externalURL(endpoint string) string {
	url := url.URL{
		Scheme: cont.ExternalProtocol,
		Host:   cont.ExternalDomain,
		Path:   endpoint,
	}
	return url.String()
}
