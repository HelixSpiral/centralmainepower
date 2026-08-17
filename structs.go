package centralmainepower

import (
	"net/http"
	"time"

	"github.com/HelixSpiral/centralmainepower/v3/load"
)

type Client struct {
	Client     *http.Client
	ReqHeaders map[string]string

	MWStatsUrl    string
	PowerStatsUrl string
	Load          load.Service
}

type Config struct {
	Client     *http.Client
	ReqHeaders map[string]string

	MWStatsUrl    string
	PowerStatsUrl string
}

type CMPPowerStats struct {
	LastUpdate   time.Time
	Total        string
	WithoutPower string
	Counties     map[string]Outage
	NoOutages    bool
}

type Outage struct {
	Total        string
	WithoutPower string
}
