package bindings

import (
	"fmt"
	"strconv"
	"strings"
)

func Bool(binding *bool) Binding {
	return &_bool{
		binding: CheckNotNil(binding),
	}
}

type _bool struct {
	binding *bool
}

func (b *_bool) Assign(str string) error {
	value, err := _boolParse(str)
	if err == nil {
		*b.binding = value
	}
	return err
}

func (b _bool) IsZero() bool {
	return !*b.binding
}

func (b *_bool) Reset() {
	*b.binding = false
}

func (b _bool) String() string {
	return _boolString(*b.binding)
}

func (b _bool) Type() string {
	return _boolType
}

func BoolFlag(binding *bool) Binding {
	return &_boolFlag{
		_bool{
			binding: CheckNotNil(binding),
		},
	}
}

type _boolFlag struct {
	_bool
}

func BoolSlice(binding *[]bool) Binding {
	return &_boolSlice{
		binding: CheckNotNil(binding),
	}
}

type _boolSlice struct {
	binding *[]bool
}

func (b *_boolSlice) Assign(str string) error {
	for part := range strings.SplitSeq(str, ",") {
		value, err := _boolParse(part)
		if err != nil {
			return err
		}
		*b.binding = append(*b.binding, value)
	}
	return nil
}

func (b _boolSlice) IsZero() bool {
	return len(*b.binding) == 0
}

func (b *_boolSlice) Reset() {
	*b.binding = nil
}

func (b _boolSlice) String() string {
	s := ""
	for i, v := range *b.binding {
		if i > 0 {
			s += ", "
		}
		s += _boolString(v)
	}
	return fmt.Sprintf("[%s]", s)
}

func (b _boolSlice) Type() string {
	return _boolType + "..."
}

const _boolType = "bool"

func _boolParse(str string) (bool, error) {
	str = strings.TrimSpace(str)
	value, err := strconv.ParseBool(str)
	if err != nil {
		err = fmt.Errorf("bad bool value %q", str)
	}
	return value, err
}

func _boolString(value bool) string {
	return fmt.Sprintf("%v", value)
}
