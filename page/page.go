package page

import (
	"bytes"
	"html/template"
)

type Page[T any] struct {
	Template string
	Data     T
	Mounted  bool
}

func NewPage[T any]() *Page[T] {
	page := Page[T]{Mounted: false}
	return &page
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
	p.Mounted = true

	return nil
}
