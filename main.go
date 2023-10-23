package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/a-lafon/go-doc-serve/utils"
)

func main() {
	fmt.Println("Starting the program")

	rootDirFlag := flag.String("d", ".", "Root directory contaning documentation")

	flag.Parse()

	rootDir, filepathError := filepath.Abs(*rootDirFlag)

	if filepathError != nil {
		log.Fatalln(filepathError)
	}

	markdownFiles, err := utils.GetMarkdownFiles(rootDir)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(markdownFiles)

	fmt.Println("End of program")
}
