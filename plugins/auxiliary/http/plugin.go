package main

import (
	"github.com/deathsgun/wiesel/pkg/api"
	"github.com/deathsgun/wiesel/pkg/api/module"
)

var Name = "http"
var Category = module.CategoryAuxiliary
var Author = "lbuening"
var Description = "HTTP auxiliary module"

func OnLoad() {
	api.RegisterModule(GetHeader)
}
