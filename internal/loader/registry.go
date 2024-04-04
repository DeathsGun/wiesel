package loader

import (
	"fmt"
	"github.com/deathsgun/wiesel/pkg/api"
	"github.com/deathsgun/wiesel/pkg/api/module"
	"github.com/deathsgun/wiesel/pkg/api/netx"
	"reflect"
	"strconv"
	"strings"
)

var CurrentModule string
var Modules = make(map[string]module.Module)
var Options = make(map[string][]*module.Option)

func init() {
	api.RegisterModule = RegisterModule
}

func RegisterModule(m module.Module) {
	options, err := extractOptions(m)
	if err != nil {
		fmt.Println(err)
		return
	}

	Modules[m.Category().String()+"/"+m.Name()] = m
	Options[m.Category().String()+"/"+m.Name()] = options
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
			option.Default = convertType(value.Type(), defaultRaw)
		}

		options = append(options, option)
	}

	return options, nil
}

func convertType(t reflect.Type, value string) any {
	switch t.Kind() {
	case reflect.String:
		return value
	case reflect.Int:
		i, _ := strconv.Atoi(value)
		return i
	case reflect.Uint:
		i, _ := strconv.ParseUint(value, 10, 64)
		return i
	case reflect.Int8:
		i, _ := strconv.ParseInt(value, 10, 8)
		return int8(i)
	case reflect.Uint8:
		i, _ := strconv.ParseUint(value, 10, 8)
		return uint8(i)
	case reflect.Int16:
		i, _ := strconv.ParseInt(value, 10, 16)
		return int16(i)
	case reflect.Uint16:
		i, _ := strconv.ParseUint(value, 10, 16)
		return uint16(i)
	case reflect.Int32:
		i, _ := strconv.ParseInt(value, 10, 32)
		return int32(i)
	case reflect.Uint32:
		i, _ := strconv.ParseUint(value, 10, 32)
		return uint32(i)
	case reflect.Int64:
		i, _ := strconv.ParseInt(value, 10, 64)
		return i
	case reflect.Uint64:
		i, _ := strconv.ParseUint(value, 10, 64)
		return i
	case reflect.Float32:
		f, _ := strconv.ParseFloat(value, 32)
		return float32(f)
	case reflect.Float64:
		f, _ := strconv.ParseFloat(value, 64)
		return f
	case reflect.Bool:
		b, _ := strconv.ParseBool(value)
		return b
	default:
		break
	}
	if t == reflect.TypeOf(&netx.HostList{}) {
		return netx.ParseHosts(value)
	}
	if t == reflect.TypeOf(netx.HostList{}) {
		return *netx.ParseHosts(value)
	}
	if t == reflect.TypeFor[[]string]() {
		return strings.Split(value, ",")
	}
	return nil
}
