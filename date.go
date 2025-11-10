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
	return _dateString(*t.binding)
}

func (t _date) Type() string {
	return _dateType
}

func DateSlice(binding *[]time.Time) Binding {
	return &_dateSlice{
		binding: CheckNotNil(binding),
	}
}

type _dateSlice struct {
	binding *[]time.Time
}

func (b *_dateSlice) Assign(str string) error {
	for part := range strings.SplitSeq(str, ",") {
		value, err := _dateParse(part)
		if err != nil {
			return err
		}
		*b.binding = append(*b.binding, value)
	}
	return nil
}

func (b _dateSlice) IsZero() bool {
	return len(*b.binding) == 0
}

func (b *_dateSlice) Reset() {
	*b.binding = nil
}

func (b _dateSlice) String() string {
	s := ""
	for i, v := range *b.binding {
		if i > 0 {
			s += ", "
		}
		s += _dateString(v)
	}
	return fmt.Sprintf("[%s]", s)
}

func (b _dateSlice) Type() string {
	return _dateType + "..."
}

const (
	_dateFormat = "2006-01-02"
	_dateType   = "date"
)

func _dateParse(str string) (time.Time, error) {
	str = strings.TrimSpace(str)
	value, err := time.Parse(_dateFormat, str)
	if err != nil {
		err = fmt.Errorf("bad date value %q", str)
	}
	return value, err
}

func _dateString(value time.Time) string {
	return value.Format(_dateFormat)
}
