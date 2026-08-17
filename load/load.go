package load

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

// cmpMWLoadURL is the URL for the text file Central Maine Power manages
const cmpMWLoadURL = "https://ecmp.cmpco.com/omni/content/cmpload.txt"

type Service struct {
	Client     *http.Client
	ReqHeaders map[string]string
	url        string
}

type Reading struct {
	Timestamp time.Time
	Load      float64
}

// New takes a client and any request headers, and returns a service
func New(client *http.Client, reqHeaders map[string]string) Service {
	return Service{
		Client:     client,
		ReqHeaders: reqHeaders,

		url: cmpMWLoadURL,
	}
}

// Latest fetches the latest text file from central Maine power and returns a reading
func (s *Service) Latest() (Reading, error) {
	var reading Reading

	req, err := http.NewRequest("GET", s.url, nil)
	if err != nil {
		return reading, fmt.Errorf("erorr creating request: %w", err)
	}

	// Set any headers provided
	for k, v := range s.ReqHeaders {
		req.Header.Set(k, v)
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return reading, fmt.Errorf("error in http GET: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return reading, fmt.Errorf("error reading http response body: %w", err)
	}
	defer resp.Body.Close()

	regTimestamp := regexp.MustCompile(`DateTime:\s*(.+)\r?\nCurrent\sInstantaneous\sNetwork\sLoad\s\(MW\):\s*([0-9]+(?:\.[0-9]+)?)`)

	loadInfo := regTimestamp.FindStringSubmatch(string(body))

	if len(loadInfo) < 3 {
		return reading, fmt.Errorf("error parsing load file: %+v", loadInfo)
	}

	reading.Timestamp, err = time.Parse("Mon Jan 2 15:04:05 MST 2006", loadInfo[1])
	if err != nil {
		return reading, fmt.Errorf("error parsing load file timestamp: %w", err)
	}

	reading.Load, err = strconv.ParseFloat(loadInfo[2], 64)
	if err != nil {
		return reading, fmt.Errorf("error parsing load file load: %w", err)
	}

	return reading, nil
}
