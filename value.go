package framework

import (
	"github.com/wlMalk/mirango/framework"
)

type value struct {
	value interface{}
}

func NewValue(value interface{}) Value {
	return &value{
		value: value,
	}
}

func (v *value) Value() interface{} {
	return v.value
}

func (v *value) String() string {
	if v.value != nil {
		rv, ok := v.value.(string)
		if ok {
			return rv
		}
	}
	return ""
}

func (v *value) Bool() bool {
	if v.value != nil {
		rv, ok := v.value.(bool)
		if ok {
			return rv
		}
	}
	return false
}

func (v *value) Int() int {
	if v.value != nil {
		rv, ok := v.value.(int)
		if ok {
			return rv
		}
	}
	return 0
}

func (v *value) Int64() int64 {
	if v.value != nil {
		rv, ok := v.value.(int64)
		if ok {
			return rv
		}
	}
	return 0
}

func (v *value) Float() float32 {
	if v.value != nil {
		rv, ok := v.value.(float32)
		if ok {
			return rv
		}
	}
	return 0
}

func (v *value) Float64() float64 {
	if v.value != nil {
		rv, ok := v.value.(float64)
		if ok {
			return rv
		}
	}
	return 0
}
