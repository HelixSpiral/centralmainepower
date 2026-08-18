package outage

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// cmpOutageURL is the URL for the central maine power outage site
const cmpOutageURL = "https://ecmp.cmpco.com/OutageReports/CMP.html"

// New takes a client and any request headers, and returns a Client
func New(client *http.Client, reqHeaders http.Header) Client {
	return Client{
		client:     client,
		reqHeaders: reqHeaders,

		url: cmpOutageURL,
	}
}

// Latest fetches the latest outage information and returns a Report
func (c *Client) Latest() (Report, error) {
	var report Report
	report.Counties = make(map[string]Outage)

	loc, err := time.LoadLocation("EST")
	if err != nil {
		return report, fmt.Errorf("error loading time information: %w", err)
	}

	regTotals := regexp.MustCompile("Total</th><th>([0-9,]+)</th><th>([0-9,]+)</th>")
	counties := regexp.MustCompile(`([a-zA-Z]+\.html)'>([a-zA-Z]+)</a>.+?([0-9,]+)</t.+?([0-9,]+)</t`)
	updatedAt := regexp.MustCompile("Update: ([^<]+)")

	req, err := http.NewRequest("GET", c.url, nil)
	if err != nil {
		return report, fmt.Errorf("erorr creating request: %w", err)
	}

	// Set any headers provided
	req.Header = c.reqHeaders.Clone()

	resp, err := c.client.Do(req)
	if err != nil {
		return report, fmt.Errorf("error in http GET: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return report, fmt.Errorf("error reading http response body: %w", err)
	}
	defer resp.Body.Close()

	if strings.Contains(string(body), "No reported electricity outages are in our system.") {
		report.NoCurrentOutage = true

		return report, nil
	}

	match := regTotals.FindStringSubmatch(string(body))
	report.TotalCustomers = match[1]
	report.TotalWithoutPower = match[2]

	match2 := updatedAt.FindStringSubmatch(string(body))

	report.Timestamp, err = time.ParseInLocation("Jan 02, 2006 03:04 PM", match2[1], loc)
	if err != nil {
		return report, fmt.Errorf("erorr parsing time location: %w", err)
	}

	match3 := counties.FindAllStringSubmatch(string(body), -1)

	for _, y := range match3 {
		report.Counties[y[2]] = Outage{
			Total:        y[3],
			WithoutPower: y[4],
		}
	}

	return report, nil
}
