export PATH := $(GOPATH)/bin:$(PATH)

all: fmt build

build:
	go build .

fmt:
	go fmt ./...


test:
	go test -v .

clean:
	rm -f ./bin/plasma-backend
