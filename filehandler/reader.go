package filehandler

import (
	"os"
	"sync"
)

type ReaderContent struct {
	Path    Path
	Content string
}

type ReaderError struct {
	Path Path
	Err  error
}

type Reader struct{}

func (r *Reader) ReadMany(paths []Path) ([]ReaderContent, []ReaderError) {
	contents := make([]ReaderContent, 0)
	errors := make([]ReaderError, 0)
	contentsChan := make(chan ReaderContent, len(paths))
	errorsChan := make(chan ReaderError)
	var wg sync.WaitGroup

	for _, path := range paths {
		wg.Add(1)
		go r.readAsync(path, &wg, contentsChan, errorsChan)
	}

	go func() {
		wg.Wait()
		close(contentsChan)
		close(errorsChan)
	}()

	for content := range contentsChan {
		contents = append(contents, content)
	}

	for err := range errorsChan {
		errors = append(errors, err)
	}

	return contents, errors
}

func (r *Reader) readAsync(path Path, wg *sync.WaitGroup, c chan<- ReaderContent, errChan chan<- ReaderError) {
	defer wg.Done()
	data, err := r.read(path)

	if err != nil {
		errChan <- ReaderError{Path: path, Err: err}
		return
	}

	c <- ReaderContent{Path: path, Content: data}
}

func (r *Reader) read(path Path) (string, error) {
	data, err := os.ReadFile(string(path))

	if err != nil {
		return "", err
	}

	return string(data), nil
}
