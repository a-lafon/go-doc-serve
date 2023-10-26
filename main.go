package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/a-lafon/go-doc-serve/filehandler"
)

func main() {
	fmt.Println("Starting the program")

	rootDirFlag := flag.String("d", ".", "Root directory contaning documentation")

	flag.Parse()

	rootDir, filepathError := filepath.Abs(*rootDirFlag)

	if filepathError != nil {
		log.Fatalln("Error parsing root directory: ", filepathError)
	}

	fileLister := filehandler.Lister{}
	fileReader := filehandler.Reader{}

	fileExtension := ".md"
	markdownPaths, filesError := fileLister.GetPathsWithExtension(rootDir, fileExtension)

	println("markdownPaths", markdownPaths)

	if filesError != nil {
		log.Fatalln("Error opening:", rootDir, filesError)
	}

	filesContent, filesErrors := fileReader.ReadMany(markdownPaths)
	fmt.Println("filesContent", filesContent)
	fmt.Println("filesErrors", filesErrors)
	fmt.Println("End of program")
}
