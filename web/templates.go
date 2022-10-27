package web

import (
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"

	"website/ui"
)

func NewTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	println(fmt.Sprintf("cache: %v", cache))

	return cache, nil
}

type templateData struct {
	CurrentYear int
}
