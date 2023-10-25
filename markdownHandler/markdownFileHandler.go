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
	markdownFilesChan := make(chan MarkdownFile, len(m.Paths))
	var wg sync.WaitGroup

	for _, path := range m.Paths {
		wg.Add(1)
		go m.readFileAsync(path, &wg, markdownFilesChan)
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

func (m *MarkdownFileHandler) readFileAsync(path string, wg *sync.WaitGroup, c chan<- MarkdownFile) {
	defer wg.Done()

	data, err := m.FileReader.ReadFile(path)

	if err != nil {
		return
	}

	c <- MarkdownFile{Path: path, Content: data}
}
