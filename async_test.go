package wishful

import (
	"testing"
	"testing/quick"
)

// Manual tests

func Test_Async(t *testing.T) {
	f := func(x string) string {
		get := Async(func(x Any) Promise {
			return Promise{}.Of(x).(Promise)
		})
		a := get(x)
		b := a.Fold(
			Identity,
			Identity,
		)
		return b.(Promise).Extract().(string)
	}
	g := func(x string) string {
		return x
	}
	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
