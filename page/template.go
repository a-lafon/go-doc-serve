package page

import (
	"errors"
	"log"

	"github.com/a-lafon/go-doc-serve/filehandler"
)

const DEFAULT_TEMPLATE_PATH = "page/default.html"

type Template struct{}

func (t *Template) GetDefault(fileReader filehandler.Reader) string {
	defaultTemplate, err := fileReader.Read(DEFAULT_TEMPLATE_PATH)
	if err != nil {
		log.Fatal(errors.New("template "+DEFAULT_TEMPLATE_PATH+" not found"), err)
	}
	return defaultTemplate
}
