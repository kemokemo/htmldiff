package htmldiff

import (
	"io"

	"github.com/andybalholm/cascadia"
	"golang.org/x/net/html"
)

// ReadHeader reads html from 'r' reader, extract header,
// and writes the header string to the 'w' writer.
func ReadHeader(r io.Reader, w io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	h := cascadia.MustCompile("head").MatchFirst(doc)
	return html.Render(w, h)
}

// ReadBody reads html from 'r' reader, extract body,
// and writes the body string to the 'w' writer.
func ReadBody(r io.Reader, w io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	h := cascadia.MustCompile("body").MatchFirst(doc)
	return html.Render(w, h)
}

// ReadHeaderAndBody reads html from 'r' reader, extract header and body.
// Then, this function writes the header string to the 'wh' writer and
// writes the body string to the 'wb' writer.
func ReadHeaderAndBody(r io.Reader, wh, wb io.Writer) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}
	h := cascadia.MustCompile("head").MatchFirst(doc)
	err = html.Render(wh, h)
	if err != nil {
		return err
	}
	h = cascadia.MustCompile("body").MatchFirst(doc)
	return html.Render(wb, h)
}
