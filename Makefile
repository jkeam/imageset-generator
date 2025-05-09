.PHONY: build
build: clean format
	GOOS=linux GOARCH=amd64 go build -o imageset-generator main.go

.PHONY: buildmac
buildmac: clean format
	GOOS=darwin GOARCH=arm64 go build -o imageset-generator-macos main.go

.PHONY: run
run:
	go run main.go

.PHONY: clean
clean:
	rm -rf ./imageset-generator ./imageset-generator-macos

.PHONE: format
format:
	go fmt ./...
