package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"io/ioutil"
	"log"
	"os"
)

var respMap map[string]interface{}

var (
	addr     = flag.String("addr", "localhost:7070", "TCP address to listen to")
	compress = flag.Bool("compress", false, "Whether to enable transparent response compression")
	file     = flag.String("file", "mock.json", "Location of mock json file")
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func main() {
	flag.Parse()
	parseMockJson(*file)

	h := requestHandler
	if *compress {
		h = fasthttp.CompressHandler(h)
	}

	fmt.Println("Starting server on", *addr)
	fmt.Println("Available paths: ")
	for k := range respMap {
		fmt.Println(k)
	}

	err := fasthttp.ListenAndServe(*addr, h)
	if err != nil {
		panic(err)
	}
}

func parseMockJson(file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully Opened", file)

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &respMap)
	if err != nil {
		log.Fatal(err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if val, ok := respMap[path]; ok {
		log.Println(path)
		ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
		ctx.Response.SetStatusCode(200)
		if err := json.NewEncoder(ctx).Encode(val); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	} else {
		ctx.Error("Path not found", fasthttp.StatusInternalServerError)
	}
}
