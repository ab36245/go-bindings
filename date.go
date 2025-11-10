package bindings

import (
	"fmt"
	"strings"
	"time"
)

func Date(binding *time.Time) Binding {
	return &_date{
		binding: CheckNotNil(binding),
	}
}

type _date struct {
	binding *time.Time
}

func (t *_date) Assign(str string) error {
	value, err := _dateParse(str)
	if err == nil {
		*t.binding = value
	}
	return err
}

func (t _date) IsZero() bool {
	return t.binding.IsZero()
}

func (t *_date) Reset() {
	*t.binding = time.Time{}
}

func (t _date) String() string {
	return t.binding.Format("2006-01-02")
}

func (t _date) Type() string {
	return "date"
}

func _dateParse(str string) (time.Time, error) {
	str = strings.TrimSpace(str)
	value, err := time.Parse("2006-01-02", str)
	if err != nil {
		err = fmt.Errorf("bad date value %q", str)
	}
	return value, err
}
