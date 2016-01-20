PRODUCT := autotune
PROFILES := $(shell cd signatures && ls *.go | grep -v api | grep -v networking | sed "s/.go$$//g")

all:
	go get
	golint ./...
	go test -cover
	go tool vet **/*.go
	go build
	$(MAKE) website-assets

website-assets:
	jq -n --arg PROFILES "$(PROFILES)" '$$PROFILES | split(" ")' > website/js/profiles.json

website:
	echo "Nothin here govn'r"

website-dev:  website
	rm -rf $$GOPATH/src/github.com/anatma/anatma.co/content/$(PRODUCT)
	cp -r ./website $$GOPATH/src/github.com/anatma/anatma.co/content/$(PRODUCT)

.PHONY: website website-dev
