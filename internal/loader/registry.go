package loader

import (
	"fmt"
	"github.com/deathsgun/wiesel/pkg/api"
	"github.com/deathsgun/wiesel/pkg/api/module"
	"reflect"
	"strconv"
)

var CurrentModule string
var Modules = make(map[string]module.Module)
var Options = make(map[string][]*module.Option)

func init() {
	api.RegisterModule = RegisterModule
	if Debug {
		InitModules()
	}
}

func RegisterModule(modules ...module.Module) {
	for _, m := range modules {
		options, err := extractOptions(m)
		if err != nil {
			fmt.Println(err)
			return
		}

		Modules[m.Category().String()+"/"+m.Name()] = m
		Options[m.Category().String()+"/"+m.Name()] = options
	}
}

func extractOptions(m module.Module) ([]*module.Option, error) {
	var options []*module.Option

	typeOf := reflect.TypeOf(m).Elem()
	valueOf := reflect.ValueOf(m).Elem()
	for i := 0; i < typeOf.NumField(); i++ {
		field := typeOf.Field(i)
		value := valueOf.Field(i)

		option := &module.Option{
			Instance: &value,
		}
		option.Name = field.Tag.Get("option")
		if option.Name == "" {
			continue
		}

		option.Description = field.Tag.Get("description")
		if requiredRaw, ok := field.Tag.Lookup("required"); ok {
			option.Required, _ = strconv.ParseBool(requiredRaw)
		}
		if defaultRaw, ok := field.Tag.Lookup("default"); ok {

			err := module.ConvertType(defaultRaw, option.Instance)
			if err != nil {
				return nil, err
			}
		}

		options = append(options, option)
	}

	return options, nil
}
