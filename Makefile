.PHONY: UDP example
BINARY="UDPexample"
DOCKERNAME="udpexample"
VERSION="v1"

fmt:
	gofmt -w ./pkg ./cmd

build: clean
	@go build -o ${BINARY}

build-linux: clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

docker: build-linux
	@docker build -t ${DOCKERNAME}:${VERSION} .

clean:
	@rm -f ${BINARY}
