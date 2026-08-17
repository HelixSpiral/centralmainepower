package load

import (
	"net/http"
	"time"
)

type Service struct {
	Client     *http.Client
	ReqHeaders map[string]string
	url        string
}

type Reading struct {
	Timestamp time.Time
	Load      float64
}
