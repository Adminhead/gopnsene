package pkg

type CoreModule struct {
	APIClient
}

// Core returns a new core module
func (c *APIClient) Core() *CoreModule {
	m := *c
	m.Module = "core"
	return &CoreModule{
		m,
	}
}
