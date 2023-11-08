package parser

import (
	"strings"

	"github.com/a-lafon/go-doc-serve/filehandler"
	"golang.org/x/exp/slices"
)

type FileParser struct {
	Converter Converter
}

func (f *FileParser) ParseToHTML(files []filehandler.ReaderContent, rootDir string) ([]struct {
	Url     string
	Content string
}, error) {
	data := []struct {
		Url     string
		Content string
	}{}
	for _, file := range files {
		content, err := f.Converter.ToHTML(file.Content)
		filepathSubStr := strings.Split(string(file.Path), "/")
		docIndex := slices.Index(filepathSubStr, rootDir)
		paths := filepathSubStr[docIndex+1:]
		url := f.constructURL(paths)

		if err != nil {
			return nil, err
		}

		data = append(data, struct {
			Url     string
			Content string
		}{Url: url, Content: content})
	}

	return data, nil
}

func (f *FileParser) constructURL(paths []string) string {
	baseUrl := f.extractbaseURL(paths)
	title := f.extractTitle(paths)
	url := baseUrl + "/" + strings.ToLower(strings.Split(title, ".md")[0])

	if len(baseUrl) >= 1 {
		url = "/" + url
	}

	return url
}

func (f *FileParser) extractbaseURL(paths []string) string {
	return strings.Join(paths[:len(paths)-1], "/")
}

func (f *FileParser) extractTitle(paths []string) string {
	return strings.ToLower(strings.Split(paths[len(paths)-1], ".md")[0])
}
