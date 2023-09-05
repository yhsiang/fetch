# Fecth

A command line tool to fetch html content

## Build

`$ go build -o fetch cmd/fetch/main.go`

## Run

- Directly run`$ go run cmd/fetch/main.go`

- After compiled `$ ./fetch https://www.google.com https://autify.com`

- Display help `$ ./fetch --help`

- Display metadata `$ ./fetch --metadata https://www.google.com`

## Docker

- Build `$ docker build . -t fetch`

- Interactive mode `$ docker run -it -v ${PWD}:/data fetch`

- Run `$ docker run -v ${PWD}:/data fetch fetch --metadata https://www.google.com`
