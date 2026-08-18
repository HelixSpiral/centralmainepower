package load

import (
	"net/http"
	"time"
)

type Client struct {
	client     *http.Client
	reqHeaders http.Header
	url        string
}

type Reading struct {
	Timestamp time.Time
	Load      float64
}
