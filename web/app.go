package web

import (
	"html/template"
	"log"
)

type App struct {
	ErrorLog      *log.Logger
	InfoLog       *log.Logger
	TemplateCache map[string]*template.Template
}
