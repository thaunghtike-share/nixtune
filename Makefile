PRODUCT := autotune
all:
	go get
	golint ./...
	go test -cover
	go vet
	go build

website:
	echo "Nothing to do Govn'r"

website-dev: website
	rm -rf $$GOPATH/src/github.com/anatma/anatma.co/content/$(PRODUCT)
	cp -r ./website $$GOPATH/src/github.com/anatma/anatma.co/content/$(PRODUCT)
