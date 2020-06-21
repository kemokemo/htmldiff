package htmldiff

import (
	"io"
	"strings"
	"text/template" // todo: use html/template

	"github.com/gobuffalo/packr/v2"
)

// CreateHTML creates a html from template with the args,
// writes the html to the writer.
func CreateHTML(w io.Writer, header, body string) error {
	box := packr.New("MyBox", "templates")
	index, err := box.FindString("index.html")
	if err != nil {
		return err
	}

	page := template.New("page")
	page, err = page.Parse(index)
	if err != nil {
		return err
	}

	values := struct {
		Head string
		Body string
	}{
		Head: header,
		Body: strings.TrimLeft(strings.TrimRight(body, "\n"), "\n"),
	}

	err = page.Execute(w, values)
	if err != nil {
		return err
	}

	return nil
}
