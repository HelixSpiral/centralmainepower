package outage

import (
	"net/http"
	"time"
)

type Client struct {
	Client     *http.Client
	ReqHeaders map[string]string
	url        string
}

type Report struct {
	Timestamp         time.Time
	TotalCustomers    string
	TotalWithoutPower string
	Counties          map[string]Outage

	NoCurrentOutage bool
}

type Outage struct {
	Total        string
	WithoutPower string
}
