package generator

import "html/template"

// RendererStrategy is an interface defining a strategy for rendering HTML
type RendererStrategy interface {
	ToHTML() (template.HTML, error)
}

// Generator is a structure that generates HTML using a specified renderer strategy
type Generator struct {
	renderer RendererStrategy
}

// SetStrategy sets the rendering strategy for the Generator
func (g *Generator) SetStrategy(r RendererStrategy) {
	g.renderer = r
}

// RenderHTML generates HTML using the specified rendering strategy
func (g *Generator) RenderHTML() (template.HTML, error) {
	return g.renderer.ToHTML()
}
