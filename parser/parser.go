package parser

type Converter interface {
	ToHTML(content string) (string, error)
}

type Transformer interface {
	PathToThree(docPath, filePath string) ([]string, error)
}
