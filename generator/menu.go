package generator

import (
	"bytes"
	"html/template"
)

type Menu struct {
	Urls []string
}

func (m *Menu) ToHTML() (template.HTML, error) {
	linkTemplate := "<li><a href=\"{{.}}\">{{.}}</a></li>"
	menuTemplate := "<ul>{{range .}}" + linkTemplate + "{{end}}</ul>"

	t, err := template.New("menu").Parse(menuTemplate)
	if err != nil {
		return "", err
	}

	var resultHTMLBuffer bytes.Buffer

	if err := t.ExecuteTemplate(&resultHTMLBuffer, "menu", m.Urls); err != nil {
		return "", err
	}

	resultHTML := resultHTMLBuffer.String()

	return template.HTML(resultHTML), nil
}
