package page

import (
	"errors"
	"html/template"

	"github.com/a-lafon/go-doc-serve/generator"
)

type Default struct {
	Title   string
	Content template.HTML
	Menu    template.HTML
}

func (d *Default) Render(contents []generator.Content, menu generator.Menu, defaultTemplate string, requestURI string) (*Page[Default], error) {
	var currentPage = NewPage[Default]()

	for i := 0; i < len(contents); i++ {
		if contents[i].Url == requestURI {
			currentPage.Data = Default{Title: contents[i].Title, Content: contents[i].Content, Menu: menu.Content}
			err := currentPage.Assemble(defaultTemplate)
			if err != nil {
				return nil, err
			}
			break
		}
	}

	if !currentPage.Mounted {
		return nil, errors.New("Page is not mounted")
	}

	return currentPage, nil
}
