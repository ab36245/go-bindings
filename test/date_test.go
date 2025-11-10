package test

import (
	"testing"
	"time"

	"github.com/ab36245/go-cli/bindings"
)

var date1 = "1997-08-28"

func dateFromString(str string) time.Time {
	t, _ := time.Parse("2006-01-02", str)
	return t
}

func TestDate(t *testing.T) {
	t.Run("Assign", func(t *testing.T) {
		run := func(t *testing.T, s string, ev time.Time, ee error) {
			var av time.Time
			b := bindings.Date(&av)
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

		t.Run(date1, func(t *testing.T) {
			run(t, date1, dateFromString(date1), nil)
		})
	})

	t.Run("IsZero", func(t *testing.T) {
		run := func(t *testing.T, v time.Time, e bool) {
			b := bindings.Date(&v)
			a := b.IsZero()
			if a != e {
				report(t, a, e)
			}
		}

		t.Run("empty", func(t *testing.T) {
			run(t, time.Time{}, true)
		})

		t.Run(date1, func(t *testing.T) {
			run(t, dateFromString(date1), false)
		})
	})

	t.Run("Reset", func(t *testing.T) {
		run := func(t *testing.T, v time.Time) {
			b := bindings.Date(&v)
			b.Reset()
			e := time.Time{}
			if !v.Equal(e) {
				report(t, v, e)
			}
		}

		t.Run("empty", func(t *testing.T) {
			run(t, time.Time{})
		})
		t.Run(date1, func(t *testing.T) {
			run(t, dateFromString(date1))
		})
	})

	t.Run("String", func(t *testing.T) {
		run := func(t *testing.T, v time.Time, e string) {
			b := bindings.Date(&v)
			a := b.String()
			if a != e {
				report(t, a, e)

			}
		}

		t.Run("empty", func(t *testing.T) {
			run(t, time.Time{}, "0001-01-01")
		})
		t.Run(date1, func(t *testing.T) {
			run(t, dateFromString(date1), date1)
		})
	})

	t.Run("Type", func(t *testing.T) {
		v := time.Time{}
		a := bindings.Date(&v).Type()
		e := "date"
		if a != e {
			report(t, a, e)
		}
	})
}
