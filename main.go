package main

import (
	"flag"
	"fmt"
	"github.com/hellojukay/ghget/network"
)

var output string

func init() {
	flag.StringVar(&output, "output", "", "output filename")
	flag.Parse()
}
func main() {
	network.NewFile("")
	fmt.Printf("%s\n", output)
}
