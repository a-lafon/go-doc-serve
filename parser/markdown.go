package parser

import (
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"golang.org/x/exp/slices"
)

type Markdown struct{}

func (m *Markdown) ToHTML(content string) (string, error) {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse([]byte(content))

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return string(markdown.Render(doc, renderer)), nil
}

func (m *Markdown) PathToThree(docPath string, filePath string) ([]string, error) {
	delimeter := "/"
	substrings := strings.Split(filePath, delimeter)
	docIndex := slices.Index(substrings, docPath)
	return substrings[docIndex+1:], nil
}
