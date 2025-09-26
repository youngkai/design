package main

import "fmt"

type Downloader interface {
	Download(uri string)
}

type template struct {
	implement
	uri string
}

type implement interface {
	download()
	save()
}

func newTemplate(i implement) *template {
	return &template{
		implement: i,
	}
}

func (t *template) Download(uri string) {
	t.uri = uri
	fmt.Print("prepare downloading\n")
	t.implement.download()
	t.implement.save()
	fmt.Print("finish downloading\n")
}

func (t *template) save() {
	fmt.Print("default save\n")
}

type HttpDownloader struct {
	*template
}

func newHttpDownloader() *HttpDownloader {
	downloader := &HttpDownloader{}
	downloader.template = newTemplate(downloader)
	return downloader
}

func (h *HttpDownloader) download() {
	fmt.Printf("download %s via http\n", h.uri)
}

func (h *HttpDownloader) save() {
	fmt.Printf("http save\n")
}

type FtpDownloader struct {
	*template
}

func newFtpDownloader() *FtpDownloader {
	downloader := &FtpDownloader{}
	downloader.template = newTemplate(downloader)
	return downloader
}

func (f *FtpDownloader) download() {
	fmt.Printf("download %s via ftp\n", f.uri)
}

func main() {
	h := newHttpDownloader()
	h.Download("http://www.baidu.com")
	f := newFtpDownloader()
	f.Download("ftp://www.baidu.com")
}
