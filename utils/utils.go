package utils

import (
	"io/fs"
	"path/filepath"
)

func GetMarkdownFiles(rootDir string) ([]string, error) {
	return getFilesWithExtension(rootDir, ".md")
}

func getFilesWithExtension(dir string, ext string) ([]string, error) {
	files := make([]string, 0)

	collectMarkdown := func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ext {
			files = append(files, path)
		}

		return nil
	}

	err := filepath.WalkDir(dir, collectMarkdown)

	if err != nil {
		return nil, err
	}

	return files, nil
}
