package module

import (
	"github.com/deathsgun/wiesel/pkg/api/netx"
	"reflect"
	"strconv"
	"strings"
)

type Option struct {
	Type        string
	Name        string
	Description string
	Required    bool
	Default     any
	Instance    *reflect.Value
}

func (o *Option) Value() any {
	if o.Instance != nil && o.Instance.IsValid() {
		if o.Instance.IsZero() {
			return o.Default
		}
		return o.Instance.Interface()
	}
	return o.Default
}

func (o *Option) SetValue(value string) {
	if o.Instance != nil && o.Instance.IsValid() && o.Instance.CanSet() {
		o.Instance.Set(reflect.ValueOf(convertType(o.Instance.Type(), value)))
		return
	}
}

func (o *Option) SetRawValue(value any) {
	if o.Instance != nil && o.Instance.IsValid() && o.Instance.CanSet() {
		o.Instance.Set(reflect.ValueOf(value))
	}
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
