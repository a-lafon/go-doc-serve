package page

import (
	"bytes"
	"html/template"
)

type Default struct {
	Title   string
	Content template.HTML
	Menu    template.HTML
}

type Page[T any] struct {
	Title    string
	Url      string
	Template string
	Data     T
}

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
