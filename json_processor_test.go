package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseMockJson(t *testing.T) {

	/*
		Happy path
	*/
	if respMap != nil {
		t.Errorf("respMap should be empty")
	}
	err := ParseMockJson("mock.json")
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if respMap == nil {
		t.Errorf("respMap should not be empty")
	}

	path := make(map[string]interface{})
	entry := make(map[string]interface{})
	entry["statusCode"] = 200
	responseBody := make(map[string]interface{})
	responseBody["message"] = "Hello worlds!"
	data := make(map[string]interface{})
	data["time"] = "now"
	responseBody["data"] = data
	worlds := []string{"cross origin world", "mars world", "moon world"}
	responseBody["worlds"] = worlds
	entry["responseBody"] = responseBody
	path["/hello/worlds"] = entry

	// can't ensure status code number or string
	if fmt.Sprint(path) != fmt.Sprint(respMap) {
		t.Errorf("respMap doesn't match \n actual   :%v \n expected :%v", respMap, path)
	}

	statusCode := reflect.ValueOf(respMap["/hello/worlds"]).MapIndex(reflect.ValueOf("statusCode")).Interface()
	if reflect.TypeOf(statusCode).Kind() != reflect.Float64 {
		t.Errorf("statusCode should be float64")
	}

	/*
		Wrong file
	*/
	err = ParseMockJson("haha.json")
	if err == nil {
		t.Errorf("expected error but nil")
	}
	if err.Error() != "open haha.json: no such file or directory" {
		t.Errorf("unexpected error %v", err)
	}

	/*
		Invalid json file
	*/
	err = ParseMockJson("tests/wrong_json.json")
	if err == nil {
		t.Errorf("expected error but nil")
	}
	if err.Error() != "invalid character 'I' looking for beginning of value" {
		t.Errorf("unexpected error %v", err)
	}

}

func TestVerifyMockJson(t *testing.T) {

	/*
		Happy path
	*/
	err := VerifyMockJson()
	if err != nil {
		t.Errorf("unexpected error %v", err)
	}


	/*
		Missing status code
	*/
	err = ParseMockJson("tests/no_statuscode.json")
	err = VerifyMockJson()
	if err == nil {
		t.Errorf("expected error but nil")
	}
	if err.Error() != "Missing `statusCode` field in /hello/worlds" {
		t.Errorf("unexpected error %v", err)
	}


	/*
		Missing response body
	*/
	err = ParseMockJson("tests/no_responsebody.json")
	err = VerifyMockJson()
	if err == nil {
		t.Errorf("expected error but nil")
	}
	if err.Error() != "Missing `responseBody` field in /hello/worlds" {
		t.Errorf("unexpected error %v", err)
	}
}
