package test

import (
	"testing"

	"github.com/ab36245/go-bindings"
)

var enum1 = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}

func TestEnum(t *testing.T) {
	t.Run("Assign", func(t *testing.T) {
		run := func(t *testing.T, s string, ev int, ee string) {
			var av int
			b := bindings.Enum(enum1, &av)
			if ae := b.Assign(s); ae != nil {
				if ae.Error() != ee {
					report(t, ae, ee)
				}
				return
			}
			if av != ev {
				report(t, av, ev)
			}
		}

		try := "one"
		t.Run(try, func(t *testing.T) {
			run(t, try, 1, "")
		})

		try = "three"
		t.Run(try, func(t *testing.T) {
			run(t, try, 3, "")
		})

		try = "invalid"
		t.Run(try, func(t *testing.T) {
			run(t, try, 0, "bad enum value \"invalid\"")
		})
	})

	t.Run("IsZero", func(t *testing.T) {
		run := func(t *testing.T, v int, e bool) {
			b := bindings.Enum(enum1, &v)
			a := b.IsZero()
			if a != e {
				report(t, a, e)
			}
		}

		t.Run("(empty)", func(t *testing.T) {
			run(t, 0, true)
		})

		t.Run("three", func(t *testing.T) {
			run(t, 3, false)
		})
	})

	t.Run("Reset", func(t *testing.T) {
	})

	t.Run("String", func(t *testing.T) {
	})

	t.Run("Type", func(t *testing.T) {
		v := 42
		a := bindings.Enum(enum1, &v).Type()
		e := "enum[int]"
		if a != e {
			report(t, a, e)
		}
	})
}
