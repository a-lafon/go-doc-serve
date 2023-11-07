package generator

import (
	"html/template"
	"strings"

	"github.com/a-lafon/go-doc-serve/filehandler"
	"github.com/a-lafon/go-doc-serve/parser"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type data struct {
	Title          string
	initialContent string
	Content        template.HTML
	Three          []string
	Url            string
}

type Content struct {
	Converter parser.Converter
	data      data
}

func (c *Content) ToHTML() (template.HTML, error) {
	html, err := c.Converter.ToHTML(c.data.initialContent)

	if err != nil {
		return "", err
	}

	return template.HTML(html), nil
}

func (c *Content) SetDataFromFile(file filehandler.ReaderContent, docDir string) {
	three := c.getThreeFromPath(docDir, string(file.Path))
	title := c.getTitle(three)
	url := c.getUrl(three)
	c.data = data{Title: title, Three: three, Url: url, initialContent: file.Content}
}

func (c *Content) getThreeFromPath(docDir string, filePath string) []string {
	delimeter := "/"
	filepathSubStr := strings.Split(filePath, delimeter)
	docIndex := slices.Index(filepathSubStr, docDir)
	paths := filepathSubStr[docIndex+1:]
	filename := strings.ToLower(strings.Split(paths[len(paths)-1], ".md")[0])
	three := paths[:len(paths)-1]
	return append(three, filename)
}

func (c *Content) getTitle(three []string) string {
	filename := three[len(three)-1]
	return cases.Title(language.English, cases.Compact).String(filename)
}

func (c *Content) getUrl(three []string) string {
	separator := "/"
	return separator + strings.ToLower(strings.Join(three, separator))
}

func (c *Content) GetData() data {
	return c.data
}
