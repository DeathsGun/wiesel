package commands

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/motd"
)

var banner = &ishell.Cmd{
	Name: "banner",
	Help: "Display an awesome wiesel banner",
	Func: func(c *ishell.Context) {
		motd.PrintBanner(c)
	},
}
