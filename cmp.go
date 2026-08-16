package centralmainepower

import "net/http"

func New(config *Config) (Client, error) {
	client := &http.Client{}
	reqHeaders := map[string]string{
		"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.9999.99 Safari/537.36",
		"Accept":          "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Accept-Language": "en-US,en;q=0.9",
		"Connection":      "keep-alive",
	}

	mwStatsUrl := "https://ecmp.cmpco.com/omni/content/cmpload.txt"
	powerStatsUrl := "https://ecmp.cmpco.com/OutageReports/CMP.html"

	if config.Client != nil {
		client = config.Client
	}

	if config.ReqHeaders != nil {
		for k, v := range config.ReqHeaders {
			reqHeaders[k] = v
		}
	}

	if config.MWStatsUrl != "" {
		mwStatsUrl = config.MWStatsUrl
	}

	if config.PowerStatsUrl != "" {
		powerStatsUrl = config.PowerStatsUrl
	}

	return Client{
		Client:     client,
		ReqHeaders: reqHeaders,

		MWStatsUrl:    mwStatsUrl,
		PowerStatsUrl: powerStatsUrl,
	}, nil
}
