package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	diff "github.com/documize/html-diff"
	htmldiff "github.com/kemokemo/htmldiff/lib"
)

const (
	exitOk = iota
	exitInvalidArg
	exitInvalidHTML
	exitFailedOperation
)

var (
	out          string
	useNewHeader bool
	help         bool
	ver          bool
)

func init() {
	flag.StringVar(&out, "o", "diff.html", "output filename")
	flag.BoolVar(&useNewHeader, "nh", true, "true: use new header, false: use old header")
	flag.BoolVar(&help, "h", false, "display help")
	flag.BoolVar(&ver, "v", false, "display version")
	flag.Parse()
}

func main() {
	os.Exit(run())
}

func run() int {
	if help {
		fmt.Fprintf(os.Stdout, "Usage: htmldiff [<option>...] <old html> <new html>\n")
		flag.PrintDefaults()
		return exitOk
	}
	if ver {
		fmt.Fprintf(os.Stdout, "htmldiff version %s.%s\n", Version, Revision)
		return exitOk
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Please set args.\n Usage: htmldiff [<option>...] <old html> <new html>\n")
		flag.PrintDefaults()
		return exitInvalidArg
	}

	fOld, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open the old html: %v\n", err)
		return exitInvalidArg
	}
	defer fOld.Close()

	fNew, err := os.Open(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open the new html: %v\n", err)
		return exitInvalidArg
	}
	defer fNew.Close()

	oHeader := bytes.NewBufferString("")
	oBody := bytes.NewBufferString("")
	err = htmldiff.ReadHeaderAndBody(fOld, oHeader, oBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read header and body from the old html: %v\n", err)
		return exitInvalidHTML
	}

	nHeader := bytes.NewBufferString("")
	nBody := bytes.NewBufferString("")
	err = htmldiff.ReadHeaderAndBody(fNew, nHeader, nBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read header and body from the new html: %v\n", err)
		return exitInvalidHTML
	}

	var header *bytes.Buffer
	if useNewHeader {
		header = nHeader
	} else {
		header = oHeader
	}

	var cfg = &diff.Config{
		Granularity:  5,
		InsertedSpan: []diff.Attribute{{Key: "style", Val: "background-color: palegreen;"}},
		DeletedSpan:  []diff.Attribute{{Key: "style", Val: "background-color: lightpink;"}},
		ReplacedSpan: []diff.Attribute{{Key: "style", Val: "background-color: lightskyblue;"}},
		CleanTags:    []string{""},
	}

	res, err := cfg.HTMLdiff([]string{oBody.String(), nBody.String()})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to generate diff: %v\n", err)
		return exitFailedOperation
	}
	diffBody := res[0]

	dir := filepath.Dir(out)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to create directory to save diff file: %v\n", err)
			return exitInvalidArg
		}
	}
	fOut, err := os.Create(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create diff html file: %v\n", err)
		return exitInvalidArg
	}
	defer fOut.Close()

	err = htmldiff.CreateHTML(fOut, header.String(), diffBody)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to save the diff html: %v\n", err)
		return exitFailedOperation
	}

	return exitOk
}
