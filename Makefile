all: build run

build:
	go build .

run:
	go run texttest_fixture.go

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out