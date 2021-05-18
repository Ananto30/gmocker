# Gmocker
Run a blazing fast mock server in just seconds! 🚀

All you need is to make a json file that contains path and response mapping. See an example [here](https://github.com/Ananto30/mocker#sample-mockjson-file).

## Run
With defaults - 
```bash
./mocker
```
**Defaults: `addr=localhost:7070` , `file=mock.json`**


With custom flags - 
```bash
./mocker -addr=<YOUR_HOST_AND_PORT> -file=<MOCK_JSON_FILE_LOCATION>
```

## Sample mock.json file
```json
{
  "path": {
    "jsonBody": ""
  }
}
```
These `path`s will be matched and the json will be sent. 

Example - 
```json
{
  "/hello/worlds": {
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
```
If a request lands in the server in path `/hello/worlds` the nested json will be sent as response.

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
