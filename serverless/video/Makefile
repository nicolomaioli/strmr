.PHONY: build clean deploy remove

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/onUpload onUpload/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/onComplete onComplete/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/getByID getByID/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

remove:
	sls remove
