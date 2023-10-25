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

type FileError struct {
	Path string
	Err  error
}

func (m *MarkdownFileHandler) GetMarkdownFiles() ([]MarkdownFile, []FileError) {
	markdownFiles := make([]MarkdownFile, 0)
	errors := make([]FileError, 0)
	markdownFilesChan := make(chan MarkdownFile, len(m.Paths))
	errorChan := make(chan FileError, len(m.Paths))
	var wg sync.WaitGroup

	for _, path := range m.Paths {
		wg.Add(1)
		go m.readFileAsync(path, &wg, markdownFilesChan, errorChan)
	}

	go func() {
		wg.Wait()
		close(markdownFilesChan)
		close(errorChan)
	}()

	for mf := range markdownFilesChan {
		markdownFiles = append(markdownFiles, mf)
	}

	for err := range errorChan {
		errors = append(errors, err)
	}

	return markdownFiles, errors
}

func (m *MarkdownFileHandler) readFileAsync(path string, wg *sync.WaitGroup, c chan<- MarkdownFile, errChan chan<- FileError) {
	defer wg.Done()
	data, err := m.FileReader.ReadFile(path)

	if err != nil {
		errChan <- FileError{Path: path, Err: err}
		return
	}

	c <- MarkdownFile{Path: path, Content: data}
}
