package commands

import (
	"fmt"
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/loader"
)

var get = &ishell.Cmd{
	Name:      "get",
	Help:      "Gets the value of a context-specific variable",
	Func:      getFunc,
	Completer: getCompleter,
}

func getCompleter(args []string) []string {
	if loader.CurrentModule == "" {
		return []string{}
	}

	options := loader.Options[loader.CurrentModule]
	values := make([]string, 0)

	for _, option := range options {
		if contains(args, option.Name) {
			continue
		}
		values = append(values, option.Name)
	}

	return values
}

func contains(args []string, arg string) bool {
	for _, a := range args {
		if a == arg {
			return true
		}
	}
	return false
}

func getFunc(c *ishell.Context) {
	if loader.CurrentModule == "" {
		c.Println("No module loaded")
		return
	}

	if len(c.Args) == 0 {
		c.Println("Usage: get var1 [var2 ...]")
		c.Println()
		c.Println("The get command is used to get the value of one or more variables.")
		c.Println()
		return
	}

	options := loader.Options[loader.CurrentModule]
	for _, arg := range c.Args {
		for _, o := range options {
			if o.Name == arg {
				c.Println(o.Name, "=>", fmt.Sprintf("%#v", o.Value()))
			}
		}
	}
}
