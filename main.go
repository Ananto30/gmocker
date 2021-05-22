package main

import (
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

var (
	addr     = flag.String("addr", "localhost:7070", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
	file     = flag.String("file", "mock.json", "Location of mock json file")
)

func main() {
	flag.Parse()
	err := ParseMockJson(*file)
	if err != nil {
		log.Fatal(err)
	}
	err = VerifyMockJson()
	if err != nil {
		log.Fatal(err)
	}
	printPaths()

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	fmt.Println("Starting server on", *addr)

	err = fasthttp.ListenAndServe(*addr, h)
	if err != nil {
		panic(err)
	}
}
