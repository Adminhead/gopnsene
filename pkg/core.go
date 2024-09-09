package pkg

type CoreModule struct {
	APIClient
}

func (c *APIClient) Core() *CoreModule {
	m := *c
	m.Module = "core"
	return &CoreModule{
		m,
	}
}
