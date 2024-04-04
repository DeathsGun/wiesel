package httpx

import (
	"io"
	"net/http"
)

func Get(url string) (*http.Response, error) {
	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", RandomUserAgent())

	return http.DefaultClient.Do(req)
}

func Post(url string, body io.Reader) (*http.Response, error) {
	req, err := NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return http.DefaultClient.Do(req)
}

func Head(url string) (*http.Response, error) {
	req, err := NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", RandomUserAgent())

	return http.DefaultClient.Do(req)
}
