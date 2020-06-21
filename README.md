# htmldiff

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![CircleCI](https://circleci.com/gh/kemokemo/htmldiff.svg?style=svg)](https://circleci.com/gh/kemokemo/htmldiff) [![codecov](https://codecov.io/gh/kemokemo/htmldiff/branch/master/graph/badge.svg)](https://codecov.io/gh/kemokemo/htmldiff) [![Go Report Card](https://goreportcard.com/badge/github.com/kemokemo/htmldiff)](https://goreportcard.com/report/github.com/kemokemo/htmldiff)

This tool compares old and new versions of html and generates color-coded html in a diff format.

![htmldiff-samle](images/htmldiff-sample.png)

## Install

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/htmldiff/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH`.

## Usage

Please set the old version of html file for `--before` and the new version for `--after`.  
Then, the color-coded html files with diff format will be saved to the html file specified by `--out`.

```sh
htmldiff --before=v1/index.html --after=v2/index.html --out=v2-diff-with-v1/index.html
```

When comparing two html files, you need to choose which header to use. By default, it uses the header of the html file specified in `--after`. To use the header of `--before` specifies `--ah` as an argument, as shown below. ( `--ah` is a flag that means "use **a**fter html's  **h**eader". )

## Information for developers

### How to build

It is recommended to build with the go module.

If you don't use the `go module`, you need to `go get` the following libraries before building.

- github.com/andybalholm/cascadia
- github.com/documize/html-diff
- github.com/gobuffalo/packr/v2

I would like to take this opportunity to thank the authors of these wonderful libraries and tools. Thank you!

### Note

Tool `packr2` is used to convert template file `template/index.html` to `packrd/packed-packr.go` for inclusion in the binary. Tool `packr2` is installed as follows.

```sh
go get -u github.com/gobuffalo/packr/v2/packr2
```

For more information, please visit [the official website](https://github.com/gobuffalo/packr/tree/master/v2).

## License

[MIT](https://github.com/kemokemo/htmldiff/blob/master/LICENSE)

## Authors

[kemokemo](https://github.com/kemokemo)

