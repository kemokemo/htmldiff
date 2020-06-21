package htmldiff

import (
	"bytes"
	"testing"
)

func TestCreateHTML(t *testing.T) {
	type args struct {
		header string
		body   string
	}
	tests := []struct {
		name    string
		args    args
		wantW   string
		wantErr bool
	}{
		// Please add unique test cases.
		{"diff html", args{header: simpleHeader, body: diffBody}, diffHTML, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			if err := CreateHTML(w, tt.args.header, tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("CreateHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("CreateHTML() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
