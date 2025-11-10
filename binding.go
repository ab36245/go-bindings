package bindings

import "log"

type Binding interface {
	Assign(string) error
	IsZero() bool
	Reset()
	String() string
	Type() string
}

func CheckNotNil[T any](binding *T) *T {
	if binding == nil {
		log.Fatal("binding must not be nil")
	}
	return binding
}
