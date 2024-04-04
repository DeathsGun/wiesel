package httpx

import (
	"io"
	"net/http"
)

func NewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", RandomUserAgent())

	return req, nil
}
