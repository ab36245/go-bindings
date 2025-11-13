package bindings

import (
	"fmt"
	"strconv"
	"strings"
)

func Int(binding *int) Binding {
	return &_int{
		binding: CheckNotNil(binding),
	}
}

type _int struct {
	binding *int
}

func (b *_int) Assign(str string) error {
	value, err := _intParse(str)
	if err == nil {
		*b.binding = value
	}
	return err
}

func (b _int) IsZero() bool {
	return *b.binding == 0
}

func (b *_int) Reset() {
	*b.binding = 0
}

func (b _int) String() string {
	return _intString(*b.binding)
}

func (b _int) Type() string {
	return _intType
}

func IntFlag(binding *int) Binding {
	return &_intFlag{
		_int{
			binding: CheckNotNil(binding),
		},
	}
}

type _intFlag struct {
	_int
}

func (b *_intFlag) Update() {
	*b.binding++
}

func IntSlice(binding *[]int) Binding {
	return &_intSlice{
		binding: CheckNotNil(binding),
	}
}

type _intSlice struct {
	binding *[]int
}

func (b *_intSlice) Assign(str string) error {
	for part := range strings.SplitSeq(str, ",") {
		value, err := _intParse(part)
		if err != nil {
			return err
		}
		*b.binding = append(*b.binding, value)
	}
	return nil
}

func (b _intSlice) IsZero() bool {
	return len(*b.binding) == 0
}

func (b *_intSlice) Reset() {
	*b.binding = nil
}

func (b _intSlice) String() string {
	s := ""
	for i, v := range *b.binding {
		if i > 0 {
			s += ", "
		}
		s += _intString(v)
	}
	return fmt.Sprintf("[%s]", s)
}

func (b _intSlice) Type() string {
	return _intType + "..."
}

const _intType = "int"

func _intParse(str string) (int, error) {
	str = strings.TrimSpace(str)
	value, err := strconv.ParseInt(str, 10, 0)
	if err != nil {
		err = fmt.Errorf("bad int value %q", str)
	}
	return int(value), err
}

func _intString(value int) string {
	return fmt.Sprintf("%v", value)
}
