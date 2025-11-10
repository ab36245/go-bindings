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

func _enumParse[T any](mapping map[string]T, str string) (string, T, error) {
	str = strings.TrimSpace(str)
	for k, v := range mapping {
		if strings.EqualFold(k, str) {
			return k, v, nil
		}
	}
	return "", *new(T), fmt.Errorf("bad enum value %q", str)
}
