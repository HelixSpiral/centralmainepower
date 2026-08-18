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

	reqHeaders := map[string]string{}
	if config.ReqHeaders != nil {
		for k, v := range config.ReqHeaders {
			reqHeaders[k] = v
		}
	}

	loadClient := load.New(client, reqHeaders)
	outageClient := outage.New(client, reqHeaders)

	return Client{
		Outage: outageClient,
		Load:   loadClient,
	}
}
