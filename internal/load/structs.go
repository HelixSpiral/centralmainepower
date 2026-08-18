package load

import (
	"net/http"
	"time"
)

type Client struct {
	Client     *http.Client
	ReqHeaders map[string]string
	url        string
}

type Reading struct {
	Timestamp time.Time
	Load      float64
}
