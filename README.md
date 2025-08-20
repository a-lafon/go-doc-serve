# go-doc-serve

This project is a simple documentation server written in Go that reads Markdown files and generates a documentation website. The server parses Markdown content, converts it to HTML, and serves it with a basic navigation menu.

## Learning Project
This project was created as a learning exercise to explore Go programming language fundamentals, including:
- Package organization and modularity
- Interface design and the Strategy pattern
- Template rendering and HTTP servers
- File handling and error management

## Getting Started
Follow these steps to run the documentation server:

Clone the repository:

```bash
git clone https://github.com/a-lafon/go-doc-serve.git
```

Navigate to the project directory:

```bash
cd go-doc-serve
```

Build the project:

```bash
go build -o doc-serve
```

Run the server with the desired root directory containing documentation:

```bash
./doc-serve -d /path/to/your/documentation
```

Access the documentation server at http://localhost:8080 in your web browser.

## Command-line Options
- -d: Specify the root directory containing documentation. By default, it uses the current directory.


## Project Structure

- main.go: The main entry point of the application.
- filehandler: Handles file-related operations such as listing files and reading their contents.
- generator: Manages the generation of HTML content for documentation.
- page: Defines structures for creating pages with HTML content.
- parser: Parses Markdown content and converts it to HTML.

## Features
- Parses Markdown files and converts them to HTML for rendering.
- Generates a basic navigation menu based on the available URLs.
- Serves documentation pages with a simple HTTP server.

