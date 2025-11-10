package bindings

import (
	"fmt"
	"strings"
)

func Enum[T any](mapping map[string]T, binding *T) Binding {
	return &_enum[T]{
		binding: CheckNotNil(binding),
		key:     "",
		mapping: mapping,
	}
}

type _enum[T any] struct {
	binding *T
	key     string
	mapping map[string]T
}

func (b *_enum[T]) Assign(str string) error {
	key, value, err := _enumParse(b.mapping, str)
	if err == nil {
		*b.binding = value
		b.key = key
	}
	return err
}

func (b _enum[T]) IsZero() bool {
	return b.key == ""
}

func (b *_enum[T]) Reset() {
	*b.binding = *new(T)
	b.key = ""
}

func (b _enum[T]) String() string {
	return b.key
}

func (b _enum[T]) Type() string {
	return fmt.Sprintf("enum[%T]", *b.binding)
}

func EnumSlice[T any](mapping map[string]T, binding *[]T) Binding {
	return &_enumSlice[T]{
		binding: CheckNotNil(binding),
		keys:    nil,
		mapping: mapping,
	}
}

type _enumSlice[T any] struct {
	binding *[]T
	keys    []string
	mapping map[string]T
}

func (b *_enumSlice[T]) Assign(str string) error {
	for part := range strings.SplitSeq(str, ",") {
		key, value, err := _enumParse(b.mapping, part)
		if err != nil {
			return err
		}
		*b.binding = append(*b.binding, value)
		b.keys = append(b.keys, key)
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
	for i, k := range b.keys {
		if i > 0 {
			s += ", "
		}
		s += k
	}
	return fmt.Sprintf("[%s]", s)
}

func (b _enumSlice[T]) Type() string {
	return _enumType[T]() + "..."
}

func _enumParse[T any](mapping map[string]T, str string) (string, T, error) {
	str = strings.TrimSpace(str)
	for k, v := range mapping {
		if strings.EqualFold(k, str) {
			return k, v, nil
		}
	}
	return "", *new(T), fmt.Errorf("bad enum value %q", str)
}

func _enumType[T any]() string {
	return fmt.Sprintf("enum[%T]", *new(T))
}
