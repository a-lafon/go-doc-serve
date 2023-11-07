package generator

import "html/template"

type RendererStrategy interface {
	ToHTML() (template.HTML, error)
}

type Generator struct {
	renderer RendererStrategy
}

func (g *Generator) SetStrategy(r RendererStrategy) {
	g.renderer = r
}

func (g *Generator) RenderHTML() (template.HTML, error) {
	return g.renderer.ToHTML()
}
