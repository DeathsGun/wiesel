package commands

import (
	"fmt"
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/loader"
	"github.com/deathsgun/wiesel/pkg/api/console"
	"os"
	"text/tabwriter"
)

var options = &ishell.Cmd{
	Name: "options",
	Help: "Displays global options or for one or more modules",
	Func: optionsFunc,
}

func optionsFunc(c *ishell.Context) {
	if loader.CurrentModule == "" {
		console.Errorln(c, "No module selected. Use the 'use' command to select a module.")
		return
	}

	options := loader.Options[loader.CurrentModule]
	if len(options) == 0 {
		c.Println("No options found for module")
		return
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', 0)
	fmt.Fprintf(w, "Module options (%s):\n\n", loader.CurrentModule)
	fmt.Fprintln(w, "Name \tCurrent Setting \tRequired \tDescription")
	fmt.Fprintln(w, "---- \t--------------- \t-------- \t-----------")

	for _, o := range options {
		fmt.Fprintf(w, "%s \t%#v \t%#v \t%s\n", o.Name, o.Value(), o.Required, o.Description)
	}

	fmt.Fprintln(w)

	w.Flush()
}
