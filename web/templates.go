package web

import (
	"html/template"
	"io/fs"
	"path/filepath"

	"website/ui"
)

var functions = template.FuncMap{
	// safe is a function to output a string without sanitizing it.
	"safe": func(s string) template.HTML {
		return template.HTML(s)
	},
}

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

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}

type templateData struct {
	CurrentYear int
}
