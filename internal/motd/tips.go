package motd

import (
	"github.com/abiosoft/ishell/v2"
	"math/rand"
)

var Tips = []string{
	"Tip: Use the `show` command to display available modules.",
	"Tip: Use the `use` command to select a module.",
	"Tip: Use the `options` command to display module options.",
	"Tip: Use the `set` command to set module options.",
	"Tip: Use the `run` command to execute the module.",
}

func PrintTip(shell *ishell.Shell) {
	shell.Println(Tips[rand.Intn(len(Tips))])
}
