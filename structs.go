package centralmainepower

import (
	"net/http"

	"github.com/HelixSpiral/centralmainepower/v3/internal/load"
	"github.com/HelixSpiral/centralmainepower/v3/internal/outage"
)

type Client struct {
	Client     *http.Client
	ReqHeaders map[string]string

	MWStatsUrl    string
	PowerStatsUrl string

	Load   load.Client
	Outage outage.Client
}

type Config struct {
	Client     *http.Client
	ReqHeaders map[string]string

	MWStatsUrl    string
	PowerStatsUrl string
}
