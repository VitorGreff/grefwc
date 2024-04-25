package main

import (
	"flag"
	"gfwc/iohandler"
)

func main() {
	var (
		filepath string
		bytes    bool
		lines    bool
		words    bool
		chars    bool
	)

	flag.BoolVar(&bytes, "c", false, "count Bytes number")
	flag.BoolVar(&lines, "l", false, "count lines number")
	flag.BoolVar(&words, "w", false, "count words number")
	flag.BoolVar(&chars, "m", false, "count characters number")
	flag.Parse()

	if flag.NFlag() == 0 {
		bytes = true
		lines = true
		words = true
		chars = true
	}

	if len(flag.Args()) > 0 {
		filepath = flag.Args()[0]
		iohandler.HandleFile(filepath, bytes, lines, words, chars)
	} else {
		iohandler.HandleStdin(bytes, lines, words, chars)
	}
}
