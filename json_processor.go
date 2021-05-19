package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

var respMap map[string]interface{}

func ParseMockJson(file string) {
	jsonFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("✔️Successfully opened:", file)

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
	fmt.Println("✔️Successfully parsed:", file)
}

func verifyMockJson() {
	for k := range respMap {
		checkFieldPresent(respMap[k], k, "statusCode")
		checkFieldPresent(respMap[k], k, "responseBody")

		// TODO: Potential checks later for new features
		//checkFieldPresent(respMap[k], k, "method")
		//method := reflect.ValueOf(respMap[k]).MapIndex(reflect.ValueOf("method")).String()
		//if method == "POST" {
		//	checkFieldPresent(respMap[k], k, "requestBody")
		//}
	}
}

func checkFieldPresent(i interface{}, key, fName string) {
	dict := reflect.ValueOf(i)
	val := dict.MapIndex(reflect.ValueOf(fName))
	if val == reflect.ValueOf(nil) {
		log.Fatalf("Missing `%v` field in %v", fName, key)
	}
}

func printPaths() {
	fmt.Println("Available paths: ")
	for k := range respMap {
		fmt.Println("=>", k)
	}
}
