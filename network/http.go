package network

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/dustin/go-humanize"
)

type FileDownloader struct {
	url   string
	count string
	Total uint64
}

func NewFile(url string) FileDownloader {
	return FileDownloader{
		url: url,
	}
}

// loading the entire file into memory.
func (file FileDownloader) Download() error {
	// Get the data
	resp, err := http.Get(file.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	filename := path.Base(resp.Request.URL.Path)

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
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (file FileDownloader) Write(p []byte) (int, error) {
	return 0, nil
}

// PrintProgress prints the progress of a file write
func (wc FileDownloader) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}
