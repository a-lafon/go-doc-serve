package parser

// Converter is an interface that defines a method for converting content to HTML
type Converter interface {
	ToHTML(content string) (string, error)
}
