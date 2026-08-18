package centralmainepower

import (
	"net/http"

	"github.com/HelixSpiral/centralmainepower/v3/internal/load"
	"github.com/HelixSpiral/centralmainepower/v3/internal/outage"
)

func New(config *Config) Client {
	client := &http.Client{}

	if config.Client != nil {
		client = config.Client
	}

	var reqHeaders http.Header
	if config.ReqHeaders != nil {
		reqHeaders = config.ReqHeaders
	}

	loadClient := load.New(client, reqHeaders.Clone())
	outageClient := outage.New(client, reqHeaders.Clone())

	return Client{
		Outage: outageClient,
		Load:   loadClient,
	}
}
