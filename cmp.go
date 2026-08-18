package centralmainepower

import (
	"net/http"

	"github.com/HelixSpiral/centralmainepower/v3/internal/load"
)

func New(config *Config) (Client, error) {
	client := &http.Client{}

	powerStatsUrl := "https://ecmp.cmpco.com/OutageReports/CMP.html"

	if config.Client != nil {
		client = config.Client
	}

	reqHeaders := map[string]string{}
	if config.ReqHeaders != nil {
		for k, v := range config.ReqHeaders {
			reqHeaders[k] = v
		}
	}

	if config.PowerStatsUrl != "" {
		powerStatsUrl = config.PowerStatsUrl
	}

	loadService := load.New(client, reqHeaders)

	return Client{
		Client:     client,
		ReqHeaders: reqHeaders,

		PowerStatsUrl: powerStatsUrl,

		Load: loadService,
	}, nil
}
