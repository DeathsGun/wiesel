package main

import (
	"github.com/abiosoft/ishell/v2"
	"github.com/abiosoft/readline"
	"github.com/deathsgun/wiesel/internal/commands"
	"github.com/deathsgun/wiesel/internal/loader"
	"github.com/deathsgun/wiesel/internal/motd"
	"github.com/deathsgun/wiesel/pkg/api/console"
)

func main() {
	shell := ishell.NewWithConfig(&readline.Config{
		Prompt: "wiesel> ",
	})
	motd.PrintTip(shell)
	motd.PrintBanner(shell)

	err := loader.LoadPlugins()
	if err != nil {
		console.Errorf(shell, "Error loading plugins: %s\n", err)
		return
	}

	commands.Register(shell)

	shell.Run()
}
