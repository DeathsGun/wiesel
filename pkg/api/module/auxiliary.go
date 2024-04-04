package module

type Auxiliary struct {
	Module
	Meta    Meta
	options []Option
}

func (m *Auxiliary) Name() string {
	return m.Meta.Name
}

func (m *Auxiliary) Description() string {
	return m.Meta.Description
}

func (m *Auxiliary) Category() Category {
	return CategoryAuxiliary
}

func (m *Auxiliary) Authors() []Author {
	return m.Meta.Authors
}

func (m *Auxiliary) CheckSupported() bool {
	return m.Meta.CheckSupported
}
