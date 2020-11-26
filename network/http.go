package network

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/dariubs/percent"
	"github.com/dustin/go-humanize"
)

type FileDownloader struct {
	url   string
	count string
	Total uint64
}

func NewFile(url string) *FileDownloader {
	return &FileDownloader{
		url: url,
	}
}

// loading the entire file into memory.
func (file *FileDownloader) Download(filename string) error {
	// Get the data
	resp, err := http.Get(file.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if filename == "" {
		filename = path.Base(resp.Request.URL.Path)
	}
	size := resp.Header.Get("Content-Length")
	filesize, _ := strconv.Atoi(size)
	file.Total = uint64(filesize)
	_, params, err := mime.ParseMediaType(resp.Header.Get("Content-Disposition"))
	if params["filename"] != "" {
		filename = params["filename"]
	}
	// Create the file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	var buffer = make([]byte, 1024*1024*4)
	var sum uint64
	for {
		n, err := resp.Body.Read(buffer)
		out.Write(buffer[:n])
		sum = sum + uint64(n)
		render(filename, sum, uint64(filesize))
		if err == io.EOF {
			break
		}
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
func render(filename string, current uint64, total uint64) {
	fmt.Printf("\rDownloading %s  %s Total %s : %.2f%%", filename, humanize.Bytes(current), humanize.Bytes(total), percent.PercentOf(int(current), int(total)))
}
