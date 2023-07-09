# IAWC ( <u>I</u>n<u>a</u>ppropriate <u>W</u>ord <u>C</u>hecker )

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT) [![test-and-build](https://github.com/kemokemo/iawc/actions/workflows/test-and-build.yaml/badge.svg)](https://github.com/kemokemo/iawc/actions/workflows/test-and-build.yaml) 

This tool check and report inappropriate words in your documents.

## Usage

1. Please set words and settings in the `iawc.yaml` file.
1. Run `iawc {your documents directory path}`.

You can get result in `{found file name}: {found word}` format.

```sh
# e.g.)
iawc sample
sample/sample.txt: 名前
```

## Install

### Homebrew

```sh
brew install kemokemo/tap/iawc
```

### Scoop

First, add my scoop-bucket.

```sh
scoop bucket add kemokemo-bucket https://github.com/kemokemo/scoop-bucket.git
```

Next, install this app by running the following.

```sh
scoop install iawc
```

### Binary

Get the latest version from [the release page](https://github.com/kemokemo/iawc/releases/latest), and download the archive file for your operating system/architecture. Unpack the archive, and put the binary somewhere in your `$PATH`.

## License

[MIT](https://github.com/kemokemo/iawc/blob/main/LICENSE)

## Author

[kemokemo](https://github.com/kemokemo)

