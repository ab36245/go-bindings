package test

import (
	"testing"

	"github.com/ab36245/go-bindings"
)

func englishToInt(binding *int) bindings.Binding {
	return bindings.Enum(binding).
		Map("one", 1).
		Map("two", 2).
		Map("three", 3)
}

func TestEnum(t *testing.T) {
	t.Run("Assign", func(t *testing.T) {
		run := func(t *testing.T, s string, ev int, ee string) {
			var av int
			b := englishToInt(&av)
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
			b := englishToInt(&v)
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
		run := func(t *testing.T, v int, e string) {
			b := englishToInt(&v)
			a := b.String()
			if a != e {
				report(t, a, e)
			}
		}

		t.Run("zero", func(t *testing.T) {
			run(t, 0, "(none)")
		})

		t.Run("one", func(t *testing.T) {
			run(t, 1, "one")
		})

		t.Run("three", func(t *testing.T) {
			run(t, 3, "three")
		})
	})

	t.Run("Type", func(t *testing.T) {
		v := 42
		a := englishToInt(&v).Type()
		e := "enum[int]"
		if a != e {
			report(t, a, e)
		}
	})
}
