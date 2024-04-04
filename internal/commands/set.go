package commands

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/loader"
	"github.com/deathsgun/wiesel/pkg/api/console"
)

var set = &ishell.Cmd{
	Name:      "set",
	Help:      "Sets the value of a context-specific variable",
	Func:      setFunc,
	Completer: setCompleter,
}

func setCompleter(args []string) []string {
	if loader.CurrentModule == "" {
		return []string{}
	}

	if len(args) == 1 {
		return []string{}
	}

	var options []string
	for _, o := range loader.Options[loader.CurrentModule] {
		options = append(options, o.Name)
	}
	return options
}

func setFunc(c *ishell.Context) {
	if loader.CurrentModule == "" {
		console.Errorln(c, "No module selected. Use the 'use' command to select a module.")
		return
	}

	if len(c.Args) < 2 {
		console.Errorln(c, "Usage: set <option> <value>")
		return
	}

	option := c.Args[0]
	value := c.Args[1]

	options := loader.Options[loader.CurrentModule]

	for _, o := range options {
		if o.Name != option {
			continue
		}
		o.SetValue(value)
		c.Println("Set", option, "=>", value)
		return
	}

}
