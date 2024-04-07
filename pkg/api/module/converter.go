package module

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func ConvertType(value string, ref *reflect.Value) error {
	if ref.Kind() == reflect.Ptr && ref.IsNil() {
		ref.Set(reflect.New(ref.Type().Elem()))
	}
	unmarshaler, ok := ref.Interface().(encoding.TextUnmarshaler)
	if ok {
		return unmarshaler.UnmarshalText([]byte(value))
	}

	switch ref.Kind() {
	case reflect.String:
		ref.SetString(value)
		return nil
	case reflect.Bool:
		v, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		ref.SetBool(v)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		ref.SetInt(v)
		return nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		ref.SetUint(v)
		return nil
	case reflect.Float32, reflect.Float64:
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		ref.SetFloat(v)
		return nil
	case reflect.Slice:
		if ref.Type().Elem().Kind() != reflect.String {
			return fmt.Errorf("unsupported type %s", ref.Type().Elem().Kind())
		}
		ref.Set(reflect.ValueOf(strings.Split(value, ",")))
		return nil
	default:
		return fmt.Errorf("unsupported type %s", ref.Kind())
	}
}
