package bindings

import (
	"fmt"
	"strings"
)

type EnumBinding[T comparable] interface {
	Binding
	Map(string, T) EnumBinding[T]
}

func Enum[T comparable](binding *T) EnumBinding[T] {
	return &_enum[T]{
		binding: CheckNotNil(binding),
	}
}

type _enumMapping[T comparable] struct {
	name  string
	value T
}

type _enum[T comparable] struct {
	binding  *T
	mappings []_enumMapping[T]
}

func (b *_enum[T]) Assign(str string) error {
	value, err := _enumParse(b.mappings, str)
	if err == nil {
		*b.binding = value
	}
	return err
}

func (b _enum[T]) IsZero() bool {
	return *b.binding == *new(T)
}

func (b *_enum[T]) Reset() {
	*b.binding = *new(T)
}

func (b _enum[T]) String() string {
	return _enumString(b.mappings, *b.binding)
}

func (b _enum[T]) Type() string {
	return _enumType(b.mappings)
}

func (b *_enum[T]) Map(name string, value T) EnumBinding[T] {
	b.mappings = append(b.mappings, _enumMapping[T]{name, value})
	return b
}

func EnumSlice[T comparable](binding *[]T) EnumBinding[T] {
	return &_enumSlice[T]{
		binding: CheckNotNil(binding),
	}
}

type _enumSlice[T comparable] struct {
	binding  *[]T
	mappings []_enumMapping[T]
}

func (b *_enumSlice[T]) Assign(str string) error {
	for part := range strings.SplitSeq(str, ",") {
		value, err := _enumParse(b.mappings, part)
		if err != nil {
			return err
		}
		*b.binding = append(*b.binding, value)
	}
	return nil
}

func (b _enumSlice[T]) IsZero() bool {
	return len(*b.binding) == 0
}

func (b *_enumSlice[T]) Reset() {
	*b.binding = nil
}

func (b _enumSlice[T]) String() string {
	s := ""
	for i, v := range *b.binding {
		if i > 0 {
			s += ", "
		}
		s += _enumString(b.mappings, v)
	}
	return fmt.Sprintf("[%s]", s)
}

func (b _enumSlice[T]) Type() string {
	return _enumType(b.mappings) + "..."
}

func (b *_enumSlice[T]) Map(name string, value T) EnumBinding[T] {
	b.mappings = append(b.mappings, _enumMapping[T]{name, value})
	return b
}

func _enumParse[T comparable](mappings []_enumMapping[T], str string) (T, error) {
	str = strings.TrimSpace(str)
	for _, m := range mappings {
		if strings.EqualFold(m.name, str) {
			return m.value, nil
		}
	}
	return *new(T), fmt.Errorf("bad enum value %q", str)
}

func _enumString[T comparable](mappings []_enumMapping[T], value T) string {
	for _, m := range mappings {
		if m.value == value {
			return m.name
		}
	}
	return "(none)"
}

func _enumType[T comparable](mappings []_enumMapping[T]) string {
	// return fmt.Sprintf("enum[%T]", *new(T))
	s := ""
	for i, m := range mappings {
		if i > 0 {
			s += "|"
		}
		s += m.name
	}
	return s
}
