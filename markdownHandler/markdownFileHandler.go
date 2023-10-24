package markdownHandler

import (
	"github.com/a-lafon/go-doc-serve/fileHandler"
)

type MarkdownFileHandler struct {
	Paths      []string
	FileReader fileHandler.FileReader
}

type MarkdownFile struct {
	Path    string
	Content string
}

func (m *MarkdownFileHandler) GetMarkdownFiles() ([]MarkdownFile, error) {
	markdownFiles := make([]MarkdownFile, 0)

	for _, path := range m.Paths {
		data, err := m.FileReader.ReadFile(path)

		if err != nil {
			return nil, err
		}

		markdownFiles = append(markdownFiles, MarkdownFile{Path: path, Content: data})
	}

	return markdownFiles, nil
}
