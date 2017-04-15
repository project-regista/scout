// Package request provides a simple abstraction
// for making HTTP requests to an API
package request

import (
	"fmt"
	"io"
	"net/http"
)

// Get makes an HTTP GET request
func Get(rURL string, params map[string]string) (io.ReadCloser, error) {

	// Prepare GET request
	req, err := http.NewRequest("GET", rURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %s", err)
	}

	// Add request options
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()

	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// Send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Failed to send request to the server: %s", err)
	}
	return resp.Body, nil
}
