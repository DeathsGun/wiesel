package commands

import (
	"fmt"
	"github.com/abiosoft/ishell/v2"
	"github.com/deathsgun/wiesel/internal/loader"
	"github.com/deathsgun/wiesel/pkg/api/module"
	"os"
	"text/tabwriter"
)

var show = &ishell.Cmd{
	Name:      "show",
	Help:      "Displays modules of a given type, or all modules",
	Func:      showFunc,
	Completer: showCompleter,
}

func showCompleter(args []string) []string {
	if len(args) > 1 {
		return []string{}
	}
	return []string{"all", "exploits", "payloads", "auxiliary", "post", "plugins"}
}

func showFunc(c *ishell.Context) {
	c.ShowPrompt(false)
	defer c.ShowPrompt(true)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)

	if c.Args[0] == "exploits" || c.Args[0] == "all" {
		fmt.Fprintln(w, "Exploits")
		fmt.Fprintln(w, "========")
		fmt.Fprintln(w, "\n# \tName \tDisclosure Date \tRank \tCheck \tDescription")
		fmt.Fprintln(w, "- \t---- \t--------------- \t---- \t----- \t-----------")

		index := 1
		for _, m := range loader.Modules {
			if m.Category() != module.CategoryExploit {
				continue
			}
			fmt.Fprintf(w, "%d \t%s \t%s \t%s \t%t \t%s\n", index, m.Name(), ".", "normal", m.CheckSupported(), m.Description())
			index++
		}

		fmt.Fprintln(w)
		w.Flush()
	}

	if c.Args[0] == "payload" || c.Args[0] == "all" {
		fmt.Fprintln(w, "Payloads")
		fmt.Fprintln(w, "========")
		fmt.Fprintln(w, "\n# \tName \tDisclosure Date \tRank \tCheck \tDescription")
		fmt.Fprintln(w, "- \t---- \t--------------- \t---- \t----- \t-----------")

		index := 1
		for _, m := range loader.Modules {
			if m.Category() != module.CategoryPayload {
				continue
			}
			fmt.Fprintf(w, "%d \t%s \t%s \t%s \t%t \t%s\n", index, m.Name(), ".", "normal", m.CheckSupported(), m.Description())
			index++
		}

		fmt.Fprintln(w)
		w.Flush()
	}

	if c.Args[0] == "auxiliary" || c.Args[0] == "all" {
		fmt.Fprintln(w, "Auxiliary")
		fmt.Fprintln(w, "========")
		fmt.Fprintln(w, "\n# \tName \tDisclosure Date \tRank \tCheck \tDescription")
		fmt.Fprintln(w, "- \t---- \t--------------- \t---- \t----- \t-----------")

		index := 1
		for _, m := range loader.Modules {
			if m.Category() != module.CategoryAuxiliary {
				continue
			}
			fmt.Fprintf(w, "%d \t%s \t%s \t%s \t%t \t%s\n", index, m.Name(), ".", "normal", m.CheckSupported(), m.Description())
			index++
		}

		fmt.Fprintln(w)
		w.Flush()
	}

	if c.Args[0] == "post" || c.Args[0] == "all" {
		fmt.Fprintln(w, "Post")
		fmt.Fprintln(w, "========")
		fmt.Fprintln(w, "\n# \tName \tDisclosure Date \tRank \tCheck \tDescription")
		fmt.Fprintln(w, "- \t---- \t--------------- \t---- \t----- \t-----------")

		index := 1
		for _, m := range loader.Modules {
			if m.Category() != module.CategoryPost {
				continue
			}
			fmt.Fprintf(w, "%d \t%s \t%s \t%s \t%t \t%s\n", index, m.Name(), ".", "normal", m.CheckSupported(), m.Description())
			index++
		}

		fmt.Fprintln(w)
		w.Flush()
	}

	if c.Args[0] == "plugins" || c.Args[0] == "all" {
		fmt.Fprintln(w, "Plugins")
		fmt.Fprintln(w, "========")
		fmt.Fprintln(w, "\n# \tName \tDescription")
		fmt.Fprintln(w, "- \t---- \t---------------")

		index := 1
		for _, p := range loader.Plugins {
			fmt.Fprintf(w, "%d \t%s \t%s\n", index, p.Name, p.Description)
			index++
		}

		fmt.Fprintln(w)
		w.Flush()
	}
}
