package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/a-lafon/go-doc-serve/filehandler"
	"github.com/a-lafon/go-doc-serve/parser"
)

func main() {
	fmt.Println("Starting the program")

	rootDirFlag := flag.String("d", ".", "Root directory contaning documentation")

	flag.Parse()

	rootDir, filepathError := filepath.Abs(*rootDirFlag)
	subDir := strings.Split(rootDir, "/")
	docDir := subDir[len(subDir)-1]

	println(docDir)

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

	var markdownConverter parser.Converter = &parser.Markdown{}
	var markdownTransformer parser.Transformer = &parser.Markdown{}

	for _, file := range filesContent {
		html, _ := markdownConverter.ToHTML(file.Content)
		three, _ := markdownTransformer.PathToThree(docDir, string(file.Path))

		println(three, html)
	}

	fmt.Println("End of program")
}
