package main

import (
	"fmt"
	"log"

	"github.com/a-lafon/go-doc-serve/utils"
)

func main() {
	fmt.Println("Starting the program")

	rootDir := "/home/arnaud/Documents/doc"

	markdownFiles, err := utils.GetMarkdownFiles(rootDir)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(markdownFiles)

	fmt.Println("End of program")
}
