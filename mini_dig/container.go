package mini_dig

import "go.uber.org/dig"

// Container DI interface
type Container interface {
	RegisterWithName(constructor interface{}, name string) error
	Register(constructor interface{}, opts ...dig.ProvideOption) error
	MustRegister(constructor interface{}, opts ...dig.ProvideOption)
	Call(function interface{}, opts ...dig.InvokeOption) error
	MustCall(function interface{}, opts ...dig.InvokeOption)
}
