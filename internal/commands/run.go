package commands

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/loader"
	"github.com/deathsgun/wiesel/pkg/api/console"
)

var run = &ishell.Cmd{
	Name: "run",
	Help: "Runs a module",
	Func: runFunc,
}

func runFunc(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)
	if loader.CurrentModule == "" {
		console.Errorln(c, "No module selected. Use the 'use' command to select a module")
		return
	}

	m, ok := loader.Modules[loader.CurrentModule]
	if !ok {
		console.Errorln(c, "Module not found")
		return
	}

	options := loader.Options[loader.CurrentModule]
	for _, option := range options {
		option.SetRawValue(option.Value())
	}

	err := m.Run(c)
	if err != nil {
		console.Errorln(c, err.Error())
	}
}
