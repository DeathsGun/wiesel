package commands

import (
	"fmt"
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/loader"
	"github.com/fatih/color"
	"strings"
)

var use = &ishell.Cmd{
	Name:      "use",
	Help:      "Interact with a module by name or search term/index",
	Func:      useFunc,
	Completer: useCompleter,
}

func useCompleter(args []string) []string {
	if len(args) > 1 {
		return []string{}
	}
	var modules []string
	for name := range loader.Modules {
		modules = append(modules, name)
	}
	return modules
}

func useFunc(c *ishell.Context) {
	if len(c.Args) == 0 {
		c.Println("No module specified")
		return
	}
	m, ok := loader.Modules[c.Args[0]]
	if !ok {
		c.Println("Module not found")
		return
	}
	module := strings.TrimPrefix(c.Args[0], m.Category().String()+"/")
	c.SetPrompt(fmt.Sprintf("wiesel %s(%s)> ", m.Category(), color.RedString(module)))
	loader.CurrentModule = c.Args[0]
}
