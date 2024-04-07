package modules

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/pkg/api/console"
	"github.com/deathsgun/wiesel/pkg/api/httpx"
	"github.com/deathsgun/wiesel/pkg/api/module"
	"github.com/deathsgun/wiesel/pkg/api/netx"
	"net/http"
	"net/netip"
	"net/url"
	"sync"
)

var GetHeader = &getHeader{
	Auxiliary: module.Auxiliary{
		Meta: module.Meta{
			Name:        "scanner/http/http_header",
			Description: "This module shows HTTP Headers returned by the scanned systems.",
			Authors: []module.Author{
				{
					Name:  "lbuening",
					Email: "lbue@proton.me",
				},
			},
		},
	},
}

type getHeader struct {
	module.Auxiliary
	Method        string         `option:"HTTP_METHOD" description:"The HTTP method to use" default:"GET"`
	IgnoreHeaders []string       `option:"IGN_HEADER" description:"List of headers to ignore, separated by comma" default:"Vary,Date,Content-Length,Connection,Etag,Expires,Pragma,Accept-Ranges"`
	HostList      *netx.HostList `option:"RHOSTS" description:"The target host(s)" required:"true"`
	Port          uint16         `option:"RPORT" description:"The target port (TCP)" default:"80" required:"true"`
	TargetUri     string         `option:"TARGETURI" description:"The URI to use" default:"/"`
	Ssl           bool           `option:"SSL" description:"Negotiate SSL/TLS for outgoing connections" default:"false"`
	VerifySsl     bool           `option:"VERIFY_SSL" description:"Verify the SSL certificate" default:"true"`
}

func (h *getHeader) Run(c *ishell.Context) error {
	wg := &sync.WaitGroup{}
	for _, host := range h.HostList.Hosts {
		wg.Add(1)
		go h.getHeaders(wg, c, netip.AddrPortFrom(host, h.Port))
	}

	wg.Wait()

	return nil
}

func (h *getHeader) getHeaders(wg *sync.WaitGroup, c *ishell.Context, host netip.AddrPort) {
	defer wg.Done()

	uri := url.URL{
		Scheme: "http",
		Host:   host.String(),
		Path:   h.TargetUri,
	}
	if h.Ssl {
		uri.Scheme = "https"
	}

	req, err := httpx.NewRequest(h.Method, uri.String(), nil)
	if err != nil {
		console.ErrorComponentf(c, host.String(), "Failed to create request: %v\n", err)
		return
	}
	var resp *http.Response
	if h.Ssl {
		resp, err = httpx.SecureClient.Do(req)
	} else {
		resp, err = httpx.InsecureClient.Do(req)
	}
	if err != nil {
		// We don't care about the error here
		return
	}
	_ = resp.Body.Close()

	count := 0
outer:
	for key := range resp.Header {
		for _, header := range h.IgnoreHeaders {
			if key == header {
				continue outer
			}
		}
		value := resp.Header.Get(key)
		console.InfoComponentf(c, host.String(), "%s: %s\n", key, value)
		count++
	}
	if count == 0 {
		console.WarningComponent(c, host.String(), "all detected headers are defined in IGN_HEADER and were ignored")
	} else {
		console.InfoComponentf(c, host.String(), "detected %d headers\n", count)
	}
}
