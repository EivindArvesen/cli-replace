.PHONY: test test-with-deps coverage check-coverage

execute-binary:
	./build/cli-replace -f test/data/myInput.txt -i test/data/myValues.toml -c test/data/myApp.toml
run:
	go run ./cmd/cli-replace/ -f test/data/myInput.txt -i test/data/myValues.toml -c test/data/myApp.toml
build:
	go build -o build/cli-replace ./cmd/cli-replace/

build-mac:
	GOOS=darwin GOARCH=amd64 build -o build/cli-replace ./cmd/cli-replace

build-windows:
	GOOS=windows GOARCH=amd64 build -o build/cli-replace ./cmd/cli-replace

build-linux:
	GOOS=linux GOARCH=amd64 build -o build/cli-replace ./cmd/cli-replace

test:
	go test ./cmd/cli-replace/...

test-with-deps:
	go test all

coverage:
	- mkdir coverage > /dev/null 2>&1
	go test -cover -coverprofile=coverage/coverage.out ./...
	go tool cover -html=coverage/coverage.out -o coverage/coverage.html

check-coverage:
	open coverage/coverage.html
