.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/upload upload/main.go
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/complete complete/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose

remove:
	sls remove
