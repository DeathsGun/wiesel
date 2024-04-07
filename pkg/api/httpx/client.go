package httpx

import (
	"crypto/tls"
	"io"
	"net/http"
	"time"
)

//TODO: Global options for timeouts

var SecureClient = &http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout:   time.Second,
		ResponseHeaderTimeout: time.Second,
	},
}

var InsecureClient = &http.Client{
	Transport: &http.Transport{
		TLSHandshakeTimeout:   time.Second,
		ResponseHeaderTimeout: time.Second,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
	},
}

func Get(url string) (*http.Response, error) {
	req, err := NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", RandomUserAgent())

	return SecureClient.Do(req)
}

func Post(url string, body io.Reader) (*http.Response, error) {
	req, err := NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return SecureClient.Do(req)
}

func Head(url string) (*http.Response, error) {
	req, err := NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", RandomUserAgent())

	return SecureClient.Do(req)
}
