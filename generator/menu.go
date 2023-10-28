package generator

import "html/template"

type Menu struct {
	Content template.HTML
}

type MenuContent struct {
	url   string
	title string
}

func (m *Menu) toHTML(content []MenuContent) {
	m.Content = template.HTML("<ul><li>I m the menu<li/></ul>")
}
