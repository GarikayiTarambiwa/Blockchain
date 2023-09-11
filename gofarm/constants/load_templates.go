package constants

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func LoadTemplates() *Template {
	template := &Template{
		templates: template.Must(template.ParseGlob("repository/templates/*.html")),
	}

	return template
}
