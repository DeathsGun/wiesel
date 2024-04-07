package modules

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/pkg/api/console"
	"github.com/deathsgun/wiesel/pkg/api/httpx"
	"github.com/deathsgun/wiesel/pkg/api/module"
	"github.com/deathsgun/wiesel/pkg/api/netx"
	"io"
	"net/netip"
)

var GetVersion = &httpVersion{
	Auxiliary: module.Auxiliary{
		Meta: module.Meta{
			Name:        "scanner/http/version",
			Description: "This module attempts to determine the version of the web server.",
			Authors: []module.Author{
				{
					Name:  "Linus Büning",
					Email: "lbue@proton.me",
				},
			},
		},
	},
}

type httpVersion struct {
	module.Auxiliary
	Hosts *netx.HostList `option:"RHOSTS" description:"The target host(s)" required:"true"`
	Port  uint16         `option:"RPORT" description:"The target port (TCP)" default:"80" required:"true"`
	Ssl   bool           `option:"SSL" description:"Negotiate SSL/TLS for outgoing connections" default:"false"`
}

func (h *httpVersion) Run(c *ishell.Context) error {

	for _, host := range h.Hosts.Hosts {
		c.Printf("Checking %s\n", host.String())
		addr := netip.AddrPortFrom(host, h.Port)
		h.printVersion(c, addr)
	}

	return nil
}

func (h *httpVersion) printVersion(c *ishell.Context, addr netip.AddrPort) {
	url := netx.FromAddrPort(addr, h.Ssl)
	resp, err := httpx.Get(url.String())
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.Header.Get("X-Powered-By") != "" {
		console.InfoComponentf(c, addr.String(), "%s (%s)\n", resp.Header.Get("Server"), resp.Header.Get("X-Powered-By"))
		return
	}

	if loc, err := resp.Location(); err == nil {
		console.InfoComponentf(c, addr.String(), "%s (%d-%s)\n", resp.Header.Get("Server"), resp.StatusCode, loc.String())
		return
	}
	console.WarningComponent(c, addr.String(), "No version information found\n")
}
