package main

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"reflect"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func requestHandler(ctx *fasthttp.RequestCtx) {
	path := string(ctx.Path())
	if val, ok := respMap[path]; ok {
		log.Println(path)
		ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)

		statusCode := int(getObject(val, "statusCode").(float64))
		ctx.Response.SetStatusCode(statusCode)

		if err := json.NewEncoder(ctx).Encode(getObject(val, "responseBody")); err != nil {
			ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
		}
	} else {
		ctx.Error("Path not found", fasthttp.StatusInternalServerError)
	}
}

func getObject(i interface{}, fieldName string) interface{} {
	return reflect.ValueOf(i).MapIndex(reflect.ValueOf(fieldName)).Interface()
}

func isPostMethod(i interface{}) bool {
	method := reflect.ValueOf(i).MapIndex(reflect.ValueOf("method")).Interface().(string)
	if method == "POST" {
		return true
	}
	return false
}
