<h1 align="center">
	<img src="https://github.com/thelicato/waybackshots/blob/main/docs/logo.png?raw=true" width="400">
</h1>

<h4 align="center">Get screenshots of URLs stored in the Wayback Machine in a smart way</h4>

<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-usage">Usage</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-license">License</a> •
</p>

---

`waybackshots` is a straightforward tool designed for obtaining screenshots from archived web pages. It allows users to easily access and capture visual records of historical web content, making it an ideal solution for **pentesters**, **bug-bounty hunters** and **security researchers**.

The process to obtain the screenshots is the following:

1. A request is made to `https://web.archive.org/cdx/search/cdx?url=<url-to-screenshot>&output=json` to obtain the list of captures from the Wayback Machine
2. The items with repeated `digest` value are discarded
3. A screenshot is made for every item

In case a file has to be handled a pre-processing step is made to reduce the number of requests to the previously mentioned endpoint.

**Unfortunately, this tool is slow** due to limitations by the Wayback Machine itself; for this reason only one URL at a time is processed and there is a timeout for connection errors (that is incremented at every error and reset once the connection is established again).

The goal is to obtain a tool that meets the requirements of the community, therefore suggestions and PRs are very welcome!

## ⚡ Features

- [x] Get screenshots for a URL
- [x] Get screenshots for all URLs in a file
- [x] Set custom output dir
- [ ] Create a neat HTML or PDF report

## 📚 Usage

```
waybackshots -h
```

This will display the help for the tool

```
               _           _       _       _
 _ _ _ ___ _ _| |_ ___ ___| |_ ___| |_ ___| |_ ___
| | | | .'| | | . | .'|  _| '_|_ -|   | . |  _|_ -|
|_____|__,|_  |___|__,|___|_,_|___|_|_|___|_| |___|
          |___|

v0.2.0 - https://github.com/thelicato/waybackshots

Get screenshots of URLs stored in the Wayback Machine in a smart way

Usage:
  waybackshots [flags]

Flags:
  -f, --file string     File with URLs to screenshot
  -h, --help            help for waybackshots
  -o, --output string   Output dir (default ".")
  -u, --url string      URL to screenshot
  -v, --verbose         Verbose mode
```

## 🚀 Installation

Run the following command to install the latest version:

```
go install github.com/thelicato/waybackshots@latest
```

## 🪪 License

_waybackshots_ is made with 🖤 and released under the [MIT LICENSE](https://github.com/thelicato/waybackshots/blob/main/LICENSE).
