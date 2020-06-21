package main

import (
	"bytes"
	"flag"
	"log"
	"os"
	"path/filepath"

	diff "github.com/documize/html-diff"
	htmldiff "github.com/kemokemo/htmldiff/lib"
)

var (
	before   = flag.String("before", "", "set the path to the before html file.")
	after = flag.String("after", "", "set the path to the after html file.")
	out    = flag.String("out", "", "set the path to save the diff html file.")
    ah = flag.Bool("ah", true, "true: use after header, false: use before header")

	bPath   string
	aPath string
	oPath    string
    useAfterHeader bool
)

func init() {
	flag.Parse()
	bPath = filepath.Clean(*before)
	aPath = filepath.Clean(*after)
	oPath = filepath.Clean(*out)
    useAfterHeader = *ah
}

func main() {
	os.Exit(run())
}

func run() int {
	fb, err := os.Open(bPath)
	if err != nil {
		log.Println("failed to open base html file:", err)
		return 1
	}
	defer fb.Close()

	bHeader := bytes.NewBufferString("")
	bBody := bytes.NewBufferString("")
	err = htmldiff.ReadHeaderAndBody(fb, bHeader, bBody)
	if err != nil {
		log.Println("failed to read header and body from base html:", err)
		return 1
	}

	fa, err := os.Open(aPath)
	if err != nil {
		log.Println("failed to open target html file:", err)
		return 1
	}
	defer fa.Close()

	aHeader := bytes.NewBufferString("")
	aBody := bytes.NewBufferString("")
	err = htmldiff.ReadHeaderAndBody(fa, aHeader, aBody)
	if err != nil {
		log.Println("failed to read body from target html:", err)
		return 1
	}

    header :=bytes.NewBufferString("")
    if(useAfterHeader){
      header = aHeader
    } else {
      header = bHeader
    }

	var cfg = &diff.Config{
		Granularity:  5,
		InsertedSpan: []diff.Attribute{{Key: "style", Val: "background-color: palegreen;"}},
		DeletedSpan:  []diff.Attribute{{Key: "style", Val: "background-color: lightpink;"}},
		ReplacedSpan: []diff.Attribute{{Key: "style", Val: "background-color: lightskyblue;"}},
		CleanTags:    []string{""},
	}

	res, err := cfg.HTMLdiff([]string{bBody.String(), aBody.String()})
	if err != nil {
		log.Println("failed to execute HTMLdiff", err)
		return 1
	}
	diffBody := res[0]

    dir := filepath.Dir(oPath)
    if _, err := os.Stat(dir); os.IsNotExist(err) {
        os.MkdirAll(dir, 0777)
    }
	fo, err := os.Create(oPath)
	if err != nil {
		log.Println("failed to create target html file:", err)
		return 1
	}
	defer fo.Close()

	err = htmldiff.CreateHTML(fo, header.String(), diffBody)
	if err != nil {
		log.Println("failed to create diff html data:", err)
		return 1
	}

	return 0
}
