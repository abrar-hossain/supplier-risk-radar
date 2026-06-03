package company

import (
	"net/http"
	"time"
)

type CompanyClient struct {
	APIKey     string
	HTTPClient http.Client
}

func NewCompanyClient(apiKey string) *CompanyClient {
	return &CompanyClient{
		APIKey: apiKey,
		HTTPClient: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
