package commands

import "github.com/abiosoft/ishell/v2"

func Register(shell *ishell.Shell) {
	shell.AddCmd(banner)
	shell.AddCmd(get)
	shell.AddCmd(options)
	shell.AddCmd(run)
	shell.AddCmd(set)
	shell.AddCmd(use)
	shell.AddCmd(show)
}
