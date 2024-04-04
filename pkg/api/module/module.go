package module

import (
	"github.com/abiosoft/ishell/v2"
)

type Module interface {
	Name() string
	Description() string
	Category() Category
	Authors() []Author
	Run(c *ishell.Context) error
	CheckSupported() bool
}

type Meta struct {
	Name           string
	Description    string
	Authors        []Author
	CheckSupported bool
}

type Author struct {
	Name  string
	Email string
}
