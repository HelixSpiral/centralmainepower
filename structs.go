package cmpscraper

import (
	"net/http"
	"time"
)

type Client struct {
	Client     *http.Client
	ReqHeaders map[string]string

	MWStatsUrl    string
	PowerStatsUrl string
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
