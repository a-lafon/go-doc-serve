package generator

import "html/template"

type Menu struct{}

func (m *Menu) ToHTML() (template.HTML, error) {
	// m.Content = template.HTML("<ul><li>I m the menu<li/></ul>")
	return "<ul><li>I m the menu<li/></ul>", nil
}
