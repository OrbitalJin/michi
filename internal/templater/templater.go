package templater

import (
	"html/template"

	"github.com/OrbitalJin/michi/public"
)

type Templater struct {
	templates *template.Template
}

func New() (*Templater, error) {
	tmplFS, err := public.SubDir("assets/templates")
	if err != nil {
		return nil, err
	}

	tmpl := template.New("")

	tmpl, err = tmpl.ParseFS(tmplFS, "*.html")
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
