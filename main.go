package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/a-lafon/go-doc-serve/filehandler"
	"github.com/a-lafon/go-doc-serve/generator"
	"github.com/a-lafon/go-doc-serve/page"
	"github.com/a-lafon/go-doc-serve/parser"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	relativePath := flag.String("d", ".", "Root directory contaning documentation")
	flag.Parse()

	absolutePath, err := filepath.Abs(*relativePath)

	if err != nil {
		log.Fatalln("Error parsing directory:", err)
	}

	rootDir := getRootDir(absolutePath)

	paths, err := getMarkdownPaths(&filehandler.Lister{}, absolutePath)

	if err != nil {
		log.Fatalln("Error opening:", rootDir, err)
	}

	fileReader := &filehandler.Reader{}
	fileContents, fileErrors := fileReader.ReadMany(paths)

	if len(fileErrors) >= 1 {
		fmt.Println("Files on error:", fileErrors)
	}

	markdownParser := parser.FileParser{Converter: &parser.Markdown{}}
	htmlContents, err := markdownParser.ParseToHTML(fileContents, rootDir)

	if err != nil {
		log.Fatalln("Error converting markdown to html:", err)
	}

	pages, err := createPages(htmlContents, *fileReader)

	if err != nil {
		log.Fatalln("Error creating pages:", err)
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
}

func getRootDir(absolutePath string) string {
	subDir := strings.Split(absolutePath, "/")
	return subDir[len(subDir)-1]
}

func getMarkdownPaths(fileLister *filehandler.Lister, absolutePath string) ([]filehandler.Path, error) {
	markdownPaths, err := fileLister.GetPathsWithExtension(absolutePath, ".md")

	if err != nil {
		return nil, err
	}

	return markdownPaths, nil
}

func createPages(htmlContents []struct {
	Url     string
	Content string
},
	fileReader filehandler.Reader) ([]page.Page[page.Default], error) {
	generatorInstance := &generator.Generator{}

	menu, err := createMenu(htmlContents, generatorInstance)

	if err != nil {
		return nil, fmt.Errorf("error generating menu: %v", err)
	}

	template := &page.Template{}
	defaultTemplate := template.GetDefault(&fileReader)

	pages := make([]page.Page[page.Default], 0)

	for _, htmlContent := range htmlContents {
		content := generator.Content{Html: htmlContent.Content}

		generatorInstance.SetStrategy(&content)
		html, err := generatorInstance.RenderHTML()

		if err != nil {
			return nil, fmt.Errorf("error generating content: %v", err)
		}

		splitedUrl := strings.Split(htmlContent.Url, "/")

		title := cases.Title(language.English, cases.Compact).String(splitedUrl[len(splitedUrl)-1])

		pageDefault := page.Default{
			Title:   title,
			Content: html,
			Menu:    menu,
		}

		page := page.Page[page.Default]{
			Title: title,
			Url:   htmlContent.Url,
			Data:  pageDefault,
		}

		err = page.Assemble(defaultTemplate)

		if err != nil {
			return nil, fmt.Errorf("error assembling page: %v", err)
		}

		pages = append(pages, page)
	}

	return pages, nil
}

func createMenu(htmlContents []struct {
	Url     string
	Content string
}, generatorInstance *generator.Generator) (template.HTML, error) {
	generatorInstance.SetStrategy(&generator.Menu{Urls: extractUrls(htmlContents)})
	menu, err := generatorInstance.RenderHTML()

	if err != nil {
		return "", fmt.Errorf("error rendering menu: %v", err)
	}

	return menu, nil
}

func extractUrls(contents []struct {
	Url     string
	Content string
}) []string {
	urls := make([]string, 0)

	for _, url := range contents {
		urls = append(urls, url.Url)
	}

	return urls
}
