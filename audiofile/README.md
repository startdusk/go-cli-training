
## To generate the audiofile command line interface:
go build -o audiofile-cli cmd/cli/main.go

## To generate the audiofile API:
go build -o audiofile-api cmd/api/main.go

## Within the root of the audiofile folder, to start the API:
./audiofile-api

### NOTE
To change the default port, 8000, pass in the new port value with the `-p` flag.

## To call the audiofile command line interface:
./audiofile-cli

### demo

run api
```bash
$ go run ./cmd/api
```

upload, will return audiofile id
```bash
$ go run ./cmd/cli upload -filename ./audio/xxx.mp3
```

get
```
$ go run ./cmd/cli get -id=0ea90dbd-392b-48ec-901c-10f39ffd190f
```
