package htmldiff

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

const (
	simpleHTML = `<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	<meta http-equiv="X-UA-Compatible" content="ie=edge"/>
	<title>すごい機能仕様書</title>
</head>
<body>
	<h1>すごい機能仕様書</h1>
	とにかくすごい。
	<h2>すごい機能1</h2>
	すごい。
	<img src="./media/screen.png" alt=""/>
</body>
</html>`

	simpleHeader = `<head>
	<meta charset="UTF-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	<meta http-equiv="X-UA-Compatible" content="ie=edge"/>
	<title>すごい機能仕様書</title>
</head>`

	simpleBody = `<body>
	<h1>すごい機能仕様書</h1>
	とにかくすごい。
	<h2>すごい機能1</h2>
	すごい。
	<img src="./media/screen.png" alt=""/>

</body>`

	badHTML = `<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8"/>
	<meta wrong-param/>
</head>
<body>
	<h1>Chapter 1</h1>
	foo.
	<h2>Chapter 1.1</h2>
	bar.
	<img src="./media/screen.png" alt="">
</body>
</html>`

	repairedHeader = `<head>
	<meta charset="UTF-8"/>
	<meta wrong-param=""/>
</head>`

	diffBody = `
<h1>すごい機能仕様書</h1>
とにかくすごい。
<h2>すごい機能1</h2>
<span style="background-color: palegreen;">なんか</span>すごい。<span style="background-color: palegreen;">とってもすごい。</span>
<img src="./media/screen.png" alt=""/><span style="background-color: palegreen;">
</span><h2><span style="background-color: palegreen;">すごい機能2</span></h2><span style="background-color: palegreen;">
こっちもすごい。</span>

`

	diffHTML = `<!DOCTYPE html>
<html lang="ja">
<head>
	<meta charset="UTF-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
	<meta http-equiv="X-UA-Compatible" content="ie=edge"/>
	<title>すごい機能仕様書</title>
</head>
<body>
<h1>すごい機能仕様書</h1>
とにかくすごい。
<h2>すごい機能1</h2>
<span style="background-color: palegreen;">なんか</span>すごい。<span style="background-color: palegreen;">とってもすごい。</span>
<img src="./media/screen.png" alt=""/><span style="background-color: palegreen;">
</span><h2><span style="background-color: palegreen;">すごい機能2</span></h2><span style="background-color: palegreen;">
こっちもすごい。</span>
</body>
</html>`
)

func TestReadHeader(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// Please add unique test cases.
		{"simple html", args{strings.NewReader(simpleHTML)}, simpleHeader, false},
		{"not html", args{strings.NewReader("hoge")}, "<head></head>", false},
		{"bad html", args{strings.NewReader(badHTML)}, repairedHeader, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := ReadHeader(tt.args.r, w); (err != nil) != tt.wantErr {
				t.Errorf("ReadHeader() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ReadHeader() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestReadBody(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// Please add unique test cases.
		{"simple html", args{strings.NewReader(simpleHTML)}, simpleBody, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := ReadBody(tt.args.r, w); (err != nil) != tt.wantErr {
				t.Errorf("ReadBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("ReadBody() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func TestReadHeaderAndBody(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantWh  string
		wantWb  string
		wantErr bool
	}{
		// Please add unique test cases.
		{"simple html", args{strings.NewReader(simpleHTML)}, simpleHeader, simpleBody, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := &bytes.Buffer{}
			wb := &bytes.Buffer{}
			if err := ReadHeaderAndBody(tt.args.r, wh, wb); (err != nil) != tt.wantErr {
				t.Errorf("ReadHeaderAndBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWh := wh.String(); gotWh != tt.wantWh {
				t.Errorf("ReadHeaderAndBody() = %v, want %v", gotWh, tt.wantWh)
			}
			if gotWb := wb.String(); gotWb != tt.wantWb {
				t.Errorf("ReadHeaderAndBody() = %v, want %v", gotWb, tt.wantWb)
			}
		})
	}
}
