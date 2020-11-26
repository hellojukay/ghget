package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hellojukay/ghget/network"
)

var output string
var url string
var proxy bool
var github_proxy = "https://g.ioiox.com/"

func init() {
	flag.StringVar(&output, "o", "", "output filename")
	flag.StringVar(&url, "url", "", "url")
	flag.BoolVar(&proxy, "proxy", false, "github proxy default: https://g.ioiox.com/ ")
	flag.StringVar(&github_proxy, "github-proxy", "https://g.ioiox.com/", "github proxy server")
	flag.Parse()
	if url == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if proxy {
		fmt.Println("donwloading by github proxy: https://g.ioiox.com/ ")
		url = github_proxy + url
	}
}
func main() {
	client := network.NewFile(url)
	client.Download(output)
	fmt.Print("\ndone\n")
}
