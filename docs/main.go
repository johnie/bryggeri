package main

import (
	"html/template"
	"os"

	"github.com/russross/blackfriday/v2"
)

func main() {
	// Read the content of the readme.md file one level up
	readmeContent, err := os.ReadFile("../readme.md")
	if err != nil {
		panic(err)
	}

	// Convert the Markdown content to HTML
	htmlContent := blackfriday.Run(readmeContent, blackfriday.WithExtensions(
		blackfriday.CommonExtensions|blackfriday.FencedCode|blackfriday.NoIntraEmphasis,
	))

	// Define the path to the HTML template file
	templatePath := "templates/index.html"

	// Read the HTML template from the template file
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}

	// Create an HTML template from the template content
	tmpl, err := template.New("index").Parse(string(templateContent))
	if err != nil {
		panic(err)
	}

	// Define a data structure to pass the HTML content to the template
	data := struct {
		Content template.HTML
	}{
		Content: template.HTML(htmlContent),
	}

	// Create the index.html file one level up
	outputFile, err := os.Create("../index.html")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Execute the template and write the output to the index.html file
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		panic(err)
	}
}
