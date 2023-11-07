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

func createPages() {

}

func main() {
	fmt.Println("Starting the program")

	rootDirFlag := flag.String("d", ".", "Root directory contaning documentation")

	flag.Parse()

	rootDir, filepathError := filepath.Abs(*rootDirFlag)
	subDir := strings.Split(rootDir, "/")
	docDir := subDir[len(subDir)-1]

	fmt.Println(docDir)

	if filepathError != nil {
		log.Fatalln("Error parsing root directory: ", filepathError)
	}

	fileLister := &filehandler.Lister{}
	fileReader := &filehandler.Reader{}

	template := &page.Template{}
	defaultTemplate := template.GetDefault(fileReader)

	markdownPaths, filesError := fileLister.GetPathsWithExtension(rootDir, ".md")

	if filesError != nil {
		log.Fatalln("Error opening:", rootDir, filesError)
	}

	filesContent, filesErrors := fileReader.ReadMany(markdownPaths)

	if len(filesErrors) >= 1 {
		fmt.Println("Files on error:", filesErrors)
	}

	generatorInstance := &generator.Generator{}

	// Generate default menu
	generatorInstance.SetStrategy(&generator.Menu{})
	menu, _ := generatorInstance.RenderHTML()

	pages := make([]page.Page[page.Default], 0)

	// Create pages
	for _, file := range filesContent {
		content := generator.Content{Converter: &parser.Markdown{}}
		content.SetDataFromFile(file, docDir)

		generatorInstance.SetStrategy(&content)

		html, err := generatorInstance.RenderHTML()

		if err != nil {
			log.Fatalln(err)
		}

		pageDefault := page.Default{
			Title:   content.GetData().Title,
			Content: html,
			Menu:    menu,
		}

		page := page.Page[page.Default]{
			Title: content.GetData().Title,
			Url:   content.GetData().Url,
			Data:  pageDefault,
		}

		err = page.Assemble(defaultTemplate)

		if err != nil {
			log.Fatalln(err)
		}

		pages = append(pages, page)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		for _, page := range pages {
			if page.Url == r.RequestURI {
				fmt.Fprint(w, page.Template)
				return
			}
		}

		fmt.Fprint(w, "Error")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("End of program")
}
