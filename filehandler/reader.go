package filehandler

import (
	"os"
	"sync"
)

// ReaderContent represents the content of a file along with its path
type ReaderContent struct {
	Path    Path
	Content string
}

// ReaderError represents an error that occurred while reading a file
type ReaderError struct {
	Path Path
	Err  error
}

// Reader is a structure that reads file content and handles multiple file reads concurrently
type Reader struct{}

// Read reads the content of a file specified by its path
func (r *Reader) Read(path Path) (string, error) {
	data, err := os.ReadFile(string(path))

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// ReadMany reads the contents of multiple files specified by their paths concurrently
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

// readAsync is a helper function for asynchronous file reading
func (r *Reader) readAsync(path Path, wg *sync.WaitGroup, c chan<- ReaderContent, errChan chan<- ReaderError) {
	defer wg.Done()
	data, err := r.Read(path)

	if err != nil {
		errChan <- ReaderError{Path: path, Err: err}
		return
	}

	c <- ReaderContent{Path: path, Content: data}
}
