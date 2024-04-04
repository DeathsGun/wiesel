package module

type Payload struct {
	Module
	Meta Meta
}

func (m *Payload) Name() string {
	return m.Meta.Name
}

func (m *Payload) Description() string {
	return m.Meta.Description
}

func (m *Payload) Category() Category {
	return CategoryPayload
}

func (m *Payload) Authors() []Author {
	return m.Meta.Authors
}

func (m *Payload) CheckSupported() bool {
	return m.Meta.CheckSupported
}
