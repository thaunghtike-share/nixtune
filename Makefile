PRODUCT := opszero
VERSION := $(shell cat VERSION)

all: build

build: # deps test
	go build -ldflags "-X main.version=$(VERSION)"

deps:
	go get -u google.golang.org/grpc
	go get -a github.com/golang/protobuf/protoc-gen-go
	go get github.com/aktau/github-release
	go get -u github.com/golang/lint/golint
	# go get -u ./...

test:
	golint ./...
	go test -cover ./...
	go tool vet **/*.go

archive:
	tar cvzf $(PRODUCT)-$(VERSION).tar.gz $(PRODUCT)

github-release:
	-github-release release --user opszero --repo $(PRODUCT) --tag v$(VERSION) --name "opszero $(VERSION)"
	-github-release upload --user opszero --repo $(PRODUCT) --tag v$(VERSION) --name "$(PRODUCT)-$(shell uname)-$(shell uname -i)-${VERSION}.tar.gz" --file $(PRODUCT)-$(VERSION).tar.gz