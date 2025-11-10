package test

import (
	"testing"

	"github.com/ab36245/go-bindings"
)

func TestBool(t *testing.T) {
	t.Run("Assign", func(t *testing.T) {
		run := func(t *testing.T, s string, ev bool, ee error) {
			var av bool
			b := bindings.Bool(&av)
			if ae := b.Assign(s); ae != nil {
				if ae != ee {
					report(t, ae, ee)
				}
				return
			}
			if av != ev {
				report(t, av, ev)
			}
		}

		t.Run("f", func(t *testing.T) {
			run(t, "f", false, nil)
		})
		t.Run("false", func(t *testing.T) {
			run(t, "false", false, nil)
		})

		t.Run("t", func(t *testing.T) {
			run(t, "t", true, nil)
		})
		t.Run("true", func(t *testing.T) {
			run(t, "true", true, nil)
		})
	})

	t.Run("IsZero", func(t *testing.T) {
		run := func(t *testing.T, v bool, e bool) {
			b := bindings.Bool(&v)
			a := b.IsZero()
			if a != e {
				report(t, a, e)
			}
		}

		t.Run("false", func(t *testing.T) {
			run(t, false, true)
		})
		t.Run("true", func(t *testing.T) {
			run(t, true, false)
		})
	})

	t.Run("Reset", func(t *testing.T) {
		run := func(t *testing.T, v bool) {
			b := bindings.Bool(&v)
			b.Reset()
			e := false
			if v != e {
				report(t, v, e)
			}
		}

		t.Run("false", func(t *testing.T) {
			run(t, false)
		})
		t.Run("true", func(t *testing.T) {
			run(t, true)
		})
	})

	t.Run("String", func(t *testing.T) {
		run := func(t *testing.T, v bool, e string) {
			b := bindings.Bool(&v)
			a := b.String()
			if a != e {
				report(t, a, e)

			}
		}

		t.Run("false", func(t *testing.T) {
			run(t, false, "false")
		})
		t.Run("true", func(t *testing.T) {
			run(t, true, "true")
		})
	})

	t.Run("Type", func(t *testing.T) {
		e := "bool"
		b := false
		a := bindings.Bool(&b).Type()
		if a != e {
			report(t, a, e)
		}
	})
}
