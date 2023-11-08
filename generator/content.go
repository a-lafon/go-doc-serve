package generator

import (
	"html/template"
)

type Content struct {
	Html string
}

func (c *Content) ToHTML() (template.HTML, error) {
	return template.HTML(c.Html), nil
}
