package bindings

import (
	"fmt"
)

func String(binding *string) Binding {
	return &_string{
		binding: CheckNotNil(binding),
	}
}

type _string struct {
	binding *string
}

func (b *_string) Assign(str string) error {
	*b.binding = str
	return nil
}

func (b _string) IsZero() bool {
	return *b.binding == ""
}

func (b *_string) Reset() {
	*b.binding = ""
}

func (b _string) String() string {
	return _stringString(*b.binding)
}

func (b _string) Type() string {
	return _stringType
}

func StringSlice(binding *[]string) Binding {
	return &_stringSlice{
		binding: CheckNotNil(binding),
	}
}

type _stringSlice struct {
	binding *[]string
}

func (b *_stringSlice) Assign(str string) error {
	*b.binding = append(*b.binding, str)
	return nil
}

func (b _stringSlice) IsZero() bool {
	return len(*b.binding) == 0
}

func (b *_stringSlice) Reset() {
	*b.binding = nil
}

func (b _stringSlice) String() string {
	if len(*b.binding) == 0 {
		return "[]"
	}
	s := "[\n"
	for _, v := range *b.binding {
		s += " "
		s += _stringString(v)
		s += ",\n"
	}
	s += "]"
	return s
}

func (b _stringSlice) Type() string {
	return _stringType + "..."
}

const _stringType = "string"

func _stringString(value string) string {
	return fmt.Sprintf("%q", value)
}
