package main

import (
	"flag"
	"fmt"
	"github.com/hellojukay/ghget/network"
)

var output string
var url string

func init() {
	flag.StringVar(&output, "o", "", "output filename")
	flag.StringVar(&url, "url", "", "url")
	flag.Parse()
}
func main() {
	client := network.NewFile(url)
	client.Download()
	fmt.Printf("%s\n", output)
}
