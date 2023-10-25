package markdownHandler

import (
	"sync"

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
	markdownFilesChan := make(chan MarkdownFile)
	var wg sync.WaitGroup

	for _, path := range m.Paths {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()

			data, err := m.FileReader.ReadFile(path)

			if err != nil {
				return
			}

			markdownFilesChan <- MarkdownFile{Path: path, Content: data}
		}(path)
	}

	go func() {
		wg.Wait()
		close(markdownFilesChan)
	}()

	for mf := range markdownFilesChan {
		markdownFiles = append(markdownFiles, MarkdownFile{Path: mf.Path, Content: mf.Content})
	}

	return markdownFiles, nil
}
