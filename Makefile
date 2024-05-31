build:
	@go build -o bin/outfile

run:
	go run ./...

start: build
	bin/outfile

test:
	go test -v ./...
