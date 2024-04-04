package banners

import (
	"github.com/deathsgun/wiesel/pkg/api/console"
	"github.com/fatih/color"
)

func PermissionDenied(p console.Printer) {
	p.Println("We trust you have received the usual lecture from the local System")
	p.Println("Administrator. It usually boils down to these three things:")
	p.Println("")
	p.Println("    #1) Respect the privacy of others.")
	p.Println("    #2) Think before you type.")
	p.Println("    #3) With great power comes great responsibility.")
	p.Println("")
	p.Println("Press [Enter] to continue.")
	p.Println("[sudo] password for wiesel:", color.New(color.FgWhite).Sprint("bakvis"))

	p.Println(color.RedString("user is not in the sudoers file. This incident will be reported."))
}
