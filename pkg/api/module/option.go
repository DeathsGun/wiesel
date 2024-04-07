package module

import (
	"fmt"
	"reflect"
)

type Option struct {
	Type        string
	Name        string
	Description string
	Required    bool
	Instance    *reflect.Value
}

func (o *Option) Value() any {
	if o.Instance != nil && o.Instance.IsValid() {
		return o.Instance.Interface()
	}
	return reflect.New(o.Instance.Type()).Interface()
}

func (o *Option) SetValue(value string) {
	if o.Instance != nil && o.Instance.IsValid() && o.Instance.CanSet() {
		err := ConvertType(value, o.Instance)
		if err != nil {
			fmt.Println(err)
		}
		return
	}
}

func (o *Option) SetRawValue(value any) {
	if o.Instance != nil && o.Instance.IsValid() && o.Instance.CanSet() {
		o.Instance.Set(reflect.ValueOf(value))
	}
}

func (o *Option) IsValid() bool {
	if o.Required && o.Instance.IsZero() {
		return false
	}
	return true
}
