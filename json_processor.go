package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

var respMap map[string]interface{}

func ParseMockJson(file string) error {
	jsonFile, err := os.Open(file)
	if err != nil {
		return err
	}
	fmt.Println("✔ Successfully opened:", file)

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(jsonFile)

	byteValue, _ := ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &respMap)
	if err != nil {
		return err
	}
	fmt.Println("✔ Successfully parsed:", file)

	return nil
}

func VerifyMockJson() error {
	for k := range respMap {
		err := checkFieldPresent(respMap[k], k, "statusCode")
		if err != nil {
			return err
		}
		err = checkFieldPresent(respMap[k], k, "responseBody")
		if err != nil {
			return err
		}

		// TODO: Potential checks later for new features
		//checkFieldPresent(respMap[k], k, "method")
		//method := reflect.ValueOf(respMap[k]).MapIndex(reflect.ValueOf("method")).String()
		//if method == "POST" {
		//	checkFieldPresent(respMap[k], k, "requestBody")
		//}
	}
	return nil
}

func checkFieldPresent(i interface{}, key, fName string) error {
	dict := reflect.ValueOf(i)
	val := dict.MapIndex(reflect.ValueOf(fName))
	if val == reflect.ValueOf(nil) {
		return errors.New(fmt.Sprintf("Missing `%v` field in %v", fName, key))
	}
	return nil
}

func printPaths() {
	fmt.Println("Available paths: ")
	for k := range respMap {
		fmt.Println("=>", k)
	}
}
