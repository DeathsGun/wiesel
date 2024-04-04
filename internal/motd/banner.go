package motd

import (
	"github.com/deathsgun/wiesel/internal/motd/banners"
	"github.com/deathsgun/wiesel/pkg/api/console"
	"math/rand"
)

var Banners = []func(console.Printer){
	banners.Cowsay,
	banners.Figlet,
	banners.PermissionDenied,
}

func PrintBanner(p console.Printer) {
	printBanner := Banners[rand.Intn(len(Banners))]
	printBanner(p)
	p.Println()
}
