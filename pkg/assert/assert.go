package assert

import (
	"reflect"
	"testing"
)

type T struct {
	Testing *testing.T
}

func New(t *testing.T) *T {
	return &T{
		Testing: t,
	}
}

func (t *T) True(f func() bool) {
	if !f() {
		t.Testing.Error("must be true")
	}
}

func (t *T) Equal(a interface{}, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Testing.Error("not equal")
	}
}

func (t *T) Err(err error) {
	if err == nil {
		t.Testing.Error("must error")
	}
}

func (t *T) Zero(a interface{}) {
	switch v := a.(type) {
	case string:
		if len(v) > 0 {
			t.Testing.Error("not zero value")
		}
	case int:
		if v > 0 {
			t.Testing.Error("not zero value")
		}
	case []string:
		if len(v) > 0 {
			t.Testing.Error("not zero value")
		}
	default:
		if !reflect.ValueOf(v).IsNil() {
			t.Testing.Error("not zero value")
		}
	}
}
