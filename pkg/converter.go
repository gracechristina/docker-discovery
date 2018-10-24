package converter

import (
	"html/template"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

type Page struct {
	Content template.HTML
}

func convertMarkdown() template.HTML {
	input := []byte("#Test")
	output := blackfriday.Run(input)
	return template.HTML(output)
}
