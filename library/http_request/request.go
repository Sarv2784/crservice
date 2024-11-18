package http_request

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Get(baseURL string, queryParams map[string]string) ([]byte, error) {
	parsedUrl, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error occured when parsing url: %w", err)
	}

	params := parsedUrl.Query()
	for key, value := range queryParams {
		params.Add(key, value)
	}
	parsedUrl.RawQuery = params.Encode()

	response, err := http.Get(parsedUrl.String())
	if err != nil {
		return nil, fmt.Errorf("error making Get call: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}
	if len(body) == 0 {
		return nil, fmt.Errorf("response body is empty")
	}

	return body, nil
}
