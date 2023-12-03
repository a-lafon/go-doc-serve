package generator

import (
	"html/template"
)

// Content is a structure representing HTML content
type Content struct {
	Html string
}

// ToHTML converts the HTML content to a template.HTML type
func (c *Content) ToHTML() (template.HTML, error) {
	return template.HTML(c.Html), nil
}
