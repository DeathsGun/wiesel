package loader

import (
	"fmt"
	"github.com/deathsgun/wiesel/pkg/api/module"
	"io/fs"
	"path/filepath"
	"plugin"
)

var Plugins = make([]Plugin, 0)

type Plugin struct {
	Name        string
	Category    module.Category
	Author      string
	Description string
}

func LoadPlugins() error {
	var files []string
	err := filepath.WalkDir("plugins", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".so" && filepath.Ext(path) != ".dylib" {
			return nil
		}

		files = append(files, path)
		return nil
	})
	if err != nil {
		return err
	}

	for _, file := range files {
		p, err := plugin.Open(file)
		if err != nil {
			return err
		}
		nameSym, err := p.Lookup("Name")
		if err != nil {
			return err
		}
		name, ok := nameSym.(*string)
		if !ok {
			return fmt.Errorf("%s: Name is not a string", file)
		}

		categorySym, err := p.Lookup("Category")
		if err != nil {
			return err
		}

		category, ok := categorySym.(*module.Category)
		if !ok {
			return fmt.Errorf("%s: Category is not a api.Category", file)
		}

		authorSym, err := p.Lookup("Author")
		if err != nil {
			return err
		}

		author, ok := authorSym.(*string)
		if !ok {
			return fmt.Errorf("%s: Author is not a string", file)
		}

		descriptionSym, err := p.Lookup("Description")
		if err != nil {
			return err
		}
		description, ok := descriptionSym.(*string)
		if !ok {
			return fmt.Errorf("%s: Description is not a string", file)
		}

		onLoad, err := p.Lookup("OnLoad")
		if err != nil {
			return err
		}
		onLoadFunc, ok := onLoad.(func())
		if !ok {
			return fmt.Errorf("%s: OnLoad is not a func()", file)
		}

		onLoadFunc()

		Plugins = append(Plugins, Plugin{
			Name:        *name,
			Category:    *category,
			Description: *description,
			Author:      *author,
		})
	}

	return nil
}
