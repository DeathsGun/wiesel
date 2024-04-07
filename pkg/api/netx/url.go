package netx

import (
	"net/netip"
	"net/url"
)

func FromAddrPort(addr netip.AddrPort, ssl bool) *url.URL {
	scheme := "http"
	if ssl {
		scheme = "https"
	}
	return &url.URL{
		Scheme: scheme,
		Host:   addr.String(),
		Path:   "/",
	}
}
