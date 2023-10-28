package generator

import (
	"html/template"
	"strings"

	"github.com/a-lafon/go-doc-serve/filehandler"
	"golang.org/x/exp/slices"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Content struct {
	Title   string
	Content template.HTML
	Three   []string
	Url     string
}

func (c *Content) ToHTML(html string, file filehandler.ReaderContent, docPath string) error {
	c.Three = c.getThreeFromPath(docPath, string(file.Path))
	c.Title = c.getTitle()
	c.Content = template.HTML(html)
	c.Url = c.getUrl()
	return nil
}

func (c *Content) getThreeFromPath(docPath string, filePath string) []string {
	delimeter := "/"
	filepathSubStr := strings.Split(filePath, delimeter)
	docIndex := slices.Index(filepathSubStr, docPath)
	paths := filepathSubStr[docIndex+1:]
	filename := strings.ToLower(strings.Split(paths[len(paths)-1], ".md")[0])
	three := paths[:len(paths)-1]
	return append(three, filename)
}

func (c *Content) getTitle() string {
	filename := c.Three[len(c.Three)-1]
	return cases.Title(language.English, cases.Compact).String(filename)
}

func (c *Content) getUrl() string {
	separator := "/"
	return separator + strings.ToLower(strings.Join(c.Three, separator))
}
