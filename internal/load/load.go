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

// New takes a client and any request headers, and returns a Client
func New(client *http.Client, reqHeaders map[string]string) Client {
	return Client{
		client:     client,
		reqHeaders: reqHeaders,

		url: cmpMWLoadURL,
	}
}

// Latest fetches the latest text file from central Maine power and returns a Reading
func (c *Client) Latest() (Reading, error) {
	var reading Reading

	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return reading, fmt.Errorf("error creating request: %w", err)
	}

	// Set any headers provided
	for k, v := range c.reqHeaders {
		req.Header.Set(k, v)
	}

	resp, err := c.client.Do(req)
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
