package centralmainepower

import (
	"net/http"

	"github.com/HelixSpiral/centralmainepower/v3/internal/load"
	"github.com/HelixSpiral/centralmainepower/v3/internal/outage"
)

type Client struct {
	Load   load.Client
	Outage outage.Client
}

type Config struct {
	Client     *http.Client
	ReqHeaders http.Header

	MWStatsUrl    string
	PowerStatsUrl string
}
