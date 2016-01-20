PRODUCT := autotune
PROFILES := $(shell cd signatures && ls *.go | grep -v api | grep -v networking | grep -v doc | sed "s/.go$$//g")
VERSION := $(shell cat VERSION)

all: build

build: deps test
	go build -ldflags "-X main.version $(VERSION)"
	$(MAKE) website-assets

deps:
	go get

test:
	golint ./...
	go test -cover
	go tool vet **/*.go

archive:
	tar cvzf autotune-$(VERSION).tar.gz autotune

release: build archive
	echo $(VERSION) > website/VERSION
	git add website/VERSION
	-git commit -m "Version $(VERSION)"
	-git tag $(VERSION) && git push --tags
	s3cmd put --acl-public nautotune s3://assets.anatma.co/autotune/${VERSION}/autotune-${VERSION}.tar.gz

website-assets:
	jq -n --arg PROFILES "$(PROFILES)" '$$PROFILES | split(" ")' > website/js/profiles.json

website:
	echo "Nothin here govn'r"

website-dev:  website
	rm -rf $$GOPATH/src/github.com/anatma/anatma.co/content/$(PRODUCT)
	cp -r ./website $$GOPATH/src/github.com/anatma/anatma.co/content/$(PRODUCT)

.PHONY: website website-dev
