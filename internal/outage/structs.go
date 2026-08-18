package outage

import (
	"net/http"
	"time"
)

type Client struct {
	client     *http.Client
	reqHeaders map[string]string
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
