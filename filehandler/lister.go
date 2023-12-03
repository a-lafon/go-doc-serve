package filehandler

import (
	"io/fs"
	"path/filepath"
)

// Lister is a structure that lists files with a specified extension in a directory
type Lister struct{}

// GetPathsWithExtension returns a list of file paths with a specified extension in the given directory
func (l *Lister) GetPathsWithExtension(dir string, ext string) ([]Path, error) {
	files := make([]Path, 0)

	collectFile := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ext {
			files = append(files, Path(path))
		}

		return nil
	}

	err := filepath.WalkDir(dir, collectFile)

	if err != nil {
		return nil, err
	}

	return files, nil
}
