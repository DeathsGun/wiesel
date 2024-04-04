package banners

import (
	"github.com/deathsgun/wiesel/pkg/api/console"
)

func Figlet(p console.Printer) {
	p.Println(" _       ___                __")
	p.Println("| |     / (_)__  ________  / /")
	p.Println("| | /| / / / _ \\/ ___/ _ \\/ /")
	p.Println("| |/ |/ / /  __(__  )  __/ /")
	p.Println("|__/|__/_/\\___/____/\\___/_/")
}
