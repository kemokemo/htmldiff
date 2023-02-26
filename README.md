# htmldiff

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![test-and-build](https://github.com/kemokemo/htmldiff/actions/workflows/test-and-build.yml/badge.svg)](https://github.com/kemokemo/htmldiff/actions/workflows/test-and-build.yml)

This tool compares old and new versions of html and generates color-coded html in a diff format.

![htmldiff-samle](images/htmldiff-sample.png)

## Install

### Homebrew

```sh
brew install kemokemo/tap/htmldiff
```

### Scoop

First, add my scoop-bucket.

```sh
scoop bucket add kemokemo-bucket https://github.com/kemokemo/scoop-bucket.git
```

Next, install this app by running the following.

```sh
scoop install htmldiff
```

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/htmldiff/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH`.

## Usage

```sh
$ htmldiff -h
Usage: htmldiff [<option>...] <old html> <new html>
  -h	display help
  -nh
    	true: use new header, false: use old header (default true)
  -o string
    	output filename (default "diff.html")j
```

When comparing two html files, you need to choose which header to use. By default, it uses the new header. To use the header of the old html file, set `-nh=false` flag. ( `-nh` means "use **n**ew **h**eader". )

### Example

```sh
htmldiff -o=diff/index.html v1/index.html v2/index.html
```

If you use the header of `v1/index.html`, set `-nh=false` flag.

```sh
htmldiff -o=diff/index.html -nh=false v1/index.html v2/index.html
```

## License

[MIT](https://github.com/kemokemo/htmldiff/blob/main/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)

