package page

import (
	"errors"
	"log"

	"github.com/a-lafon/go-doc-serve/filehandler"
)

const DEFAULT_TEMPLATE_PATH = "static/html"

type Template struct{}

func (t *Template) GetDefault(fileReader *filehandler.Reader) string {
	filePath := DEFAULT_TEMPLATE_PATH + "/default.html"
	defaultTemplate, err := fileReader.Read(filehandler.Path(filePath))
	if err != nil {
		log.Fatal(errors.New("template "+filePath+" not found"), err)
	}
	return defaultTemplate
}
