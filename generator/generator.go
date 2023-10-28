package generator

import (
	"github.com/a-lafon/go-doc-serve/filehandler"
	"github.com/a-lafon/go-doc-serve/parser"
)

type Generator struct {
	Converter parser.Converter
}

func (g *Generator) HtmlContents(readerContents []filehandler.ReaderContent, docPath string) []Content {
	contents := make([]Content, 0)
	//TODO goroutines
	for _, file := range readerContents {
		html, _ := g.Converter.ToHTML(file.Content)
		content := Content{}
		_ = content.ToHTML(html, file, docPath)
		contents = append(contents, content)
	}

	return contents
}

func (g *Generator) HtmlMenu(contents []Content) Menu {
	menu := Menu{}
	menuContents := make([]MenuContent, 0)

	for _, html := range contents {
		menuContents = append(menuContents, MenuContent{url: html.Url, title: html.Title})
	}

	menu.toHTML(menuContents)

	return menu
}
