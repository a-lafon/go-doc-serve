package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/a-lafon/go-doc-serve/filehandler"
	"github.com/a-lafon/go-doc-serve/generator"
	"github.com/a-lafon/go-doc-serve/page"
	"github.com/a-lafon/go-doc-serve/parser"
)

func main() {
	fmt.Println("Starting the program")

	rootDirFlag := flag.String("d", ".", "Root directory contaning documentation")

	flag.Parse()

	rootDir, filepathError := filepath.Abs(*rootDirFlag)
	subDir := strings.Split(rootDir, "/")
	docDir := subDir[len(subDir)-1]

	if filepathError != nil {
		log.Fatalln("Error parsing root directory: ", filepathError)
	}

	fileLister := filehandler.Lister{}
	fileReader := filehandler.Reader{}

	template := page.Template{}
	defaultTemplate := template.GetDefault(fileReader)

	fileExtension := ".md"
	markdownPaths, filesError := fileLister.GetPathsWithExtension(rootDir, fileExtension)

	println("markdownPaths", markdownPaths)

	if filesError != nil {
		log.Fatalln("Error opening:", rootDir, filesError)
	}

	filesContent, filesErrors := fileReader.ReadMany(markdownPaths)
	fmt.Println("filesErrors", filesErrors)

	var markdownConverter parser.Converter = &parser.Markdown{}
	generator := generator.Generator{Converter: markdownConverter}

	htmlContents := generator.HtmlContents(filesContent, docDir)
	htmlMenu := generator.HtmlMenu(htmlContents)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		pageDefault := page.Default{}
		currentPage, _ := pageDefault.Render(htmlContents, htmlMenu, defaultTemplate, r.RequestURI)

		// if err != nil {
		// 	// log.Fatalln(err)
		// }

		fmt.Fprint(w, currentPage.Template)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("End of program")
}
