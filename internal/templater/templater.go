package templater

import (
	"embed"
	"html/template"
	"io/fs"
)

//go:embed templates/*
var embeddedTemplates embed.FS

type Templater struct {
	templates *template.Template
}

func New() (*Templater, error) {
	templatesFS, err := fs.Sub(embeddedTemplates, "templates")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("")

	tmpl, err = tmpl.ParseFS(templatesFS, "*.html")
	if err != nil {
		return nil, err
	}

	return &Templater{
		templates: tmpl,
	}, nil
}

func (t *Templater) GetHTMLTemplates() *template.Template {
	return t.templates
}
