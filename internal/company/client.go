package company

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func (c *CompanyClient) SearchCompanies(name string) ([]SearchResult, error) {
	baseURL := "https://api.company-information.service.gov.uk/search/companies"
	searchURL := baseURL + "?q=" + url.QueryEscape(name)

	req, err := http.NewRequest("GET", searchURL, nil)

	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.APIKey, "")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Print for debugging
	//fmt.Println(string(body))

	var result SearchResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
