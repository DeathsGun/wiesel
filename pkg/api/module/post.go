package module

type Post struct {
	Module
	Meta Meta
}

func (m *Post) Name() string {
	return m.Meta.Name
}

func (m *Post) Description() string {
	return m.Meta.Description
}

func (m *Post) Category() Category {
	return CategoryPost
}

func (m *Post) Authors() []Author {
	return m.Meta.Authors
}

func (m *Post) CheckSupported() bool {
	return m.Meta.CheckSupported
}
