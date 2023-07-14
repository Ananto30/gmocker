# Gmocker
[![codecov](https://codecov.io/gh/Ananto30/gmocker/branch/main/graph/badge.svg?token=ulDcbeJyak)](https://codecov.io/gh/Ananto30/gmocker)
[![Maintainability](https://api.codeclimate.com/v1/badges/8d035908151fad8489ee/maintainability)](https://codeclimate.com/github/Ananto30/mocker/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ananto30/mocker)](https://goreportcard.com/report/github.com/Ananto30/mocker)

Run a blazing fast mock server in just seconds! 🚀

All you need is to make a json file that contains path and response mapping. See an example [here](https://github.com/Ananto30/mocker#sample-mockjson-file).

*Only json is supported for now, please create issues for bugs and new features.*

## Run
With defaults - 
```bash
./mocker
```
**Defaults: `addr=localhost:7070` , `file=mock.json`**


With custom flags - 
```bash
./mocker -addr <YOUR_HOST_AND_PORT> -file <MOCK_JSON_FILE_LOCATION>
```


For windows - 
```powershell
mocker.exe -addr <YOUR_HOST_AND_PORT> -file <MOCK_JSON_FILE_LOCATION>
```

## Sample mock.json file
```
{
  "<YOUR_PATH>": {
    "statusCode": <INTEGER>,
    "responseBody": {
      <YOUR_RESPONSE_BODY> ...
    }
  }
}
```
These `path`s will be matched and the json will be sent. 

Example - 
```json
{
  "/hello/worlds": {
    "statusCode": 200,
    "responseBody": {
      "message": "Hello worlds!",
      "data" : {
        "time": "now"
      },
      "worlds": [
        "cross origin world",
        "mars world",
        "moon world"
      ]
    }
  }
}
```
If a request lands in the server in path `/hello/worlds` the json object inside `responseBody` will be sent as response.

**The request type [POST or GET] doesn't matter.**

## Build
For mac/linux - 
```bash
go mod download
go build
```

For windows -
```bash
go mod download
GOOS=windows GOARCH=amd64 go build 
```

**If the build/binary doesn't work for you, you can do this -

- Check your os and arch using this command - `go env GOOS GOARCH`
- Use the output os and arch to build the binary - `GOOS=<YOUR_OS> GOARCH=<YOUR_ARCH> go build`
