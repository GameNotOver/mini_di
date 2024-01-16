package mini_dig

import "go.uber.org/dig"

// AppContainer 基于 uber dig 实现 DI interface
type AppContainer struct {
	*dig.Container
}

func (c *AppContainer) RegisterWithName(constructor interface{}, name string) error {
	return c.Provide(constructor, dig.Name(name))
}

func (c *AppContainer) Register(constructor interface{}, opts ...dig.ProvideOption) error {
	return c.Provide(constructor, opts...)
}

func (c *AppContainer) MustRegister(constructor interface{}, opts ...dig.ProvideOption) {
	if err := c.Provide(constructor, opts...); err != nil {
		panic(err)
	}
}

func (c *AppContainer) Call(function interface{}, opts ...dig.InvokeOption) error {
	return c.Invoke(function, opts...)
}

func (c *AppContainer) MustCall(function interface{}, opts ...dig.InvokeOption) {
	if err := c.Invoke(function, opts...); err != nil {
		panic(err)
	}
}

var appContainer *AppContainer

// GetAppContainer 获取App Container单例
func GetAppContainer() Container {
	if appContainer == nil {
		appContainer = &AppContainer{
			dig.New(),
		}
	}

	return appContainer
}
