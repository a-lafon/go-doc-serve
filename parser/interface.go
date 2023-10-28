package parser

type Converter interface {
	ToHTML(content string) (string, error)
}
