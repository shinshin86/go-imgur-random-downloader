# go-imgur-random-downloader

This is [imgur](https://imgur.com) random image downloader implemented in Go



## Install

```bash
go get github.com/shinshin86/go-imgur-random-downloader
```



## Usage

If no argument is given, a single image is downloaded to the current directory.

```bash
go-imgur-random-downloader
```



If you want to download 5 image files into the "images" directory, use this.

```bash
go-imgur-random-downloader -n=5 -path=images
```



If you want to view the help

```bash
go-imgur-random-downloader -h
```