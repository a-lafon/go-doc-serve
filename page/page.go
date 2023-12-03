package page

import (
	"bytes"
	"html/template"
)

// Default represents the default structure for a web page
type Default struct {
	Title   string
	Content template.HTML
	Menu    template.HTML
}

// Page is a generic structure for a web page with templated data
type Page[T any] struct {
	Title    string
	Url      string
	Template string
	Data     T
}

// Assemble compiles the HTML template for the Page using the provided HTML template string
func (p *Page[T]) Assemble(htmlTemplate string) error {
	t, err := template.New("Template").Parse(htmlTemplate)

	if err != nil {
		return err
	}

	var buffer bytes.Buffer
	err = t.Execute(&buffer, p.Data)

	if err != nil {
		return err
	}

	p.Template = buffer.String()

	return nil
}
