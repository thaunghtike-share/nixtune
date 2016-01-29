PRODUCT := autotune
PROFILES := $(shell cd signatures && ls *.go | grep -v signature | grep -v networking | grep -v utils | grep -v doc | sed "s/.go$$//g")
VERSION := $(shell cat VERSION)
WEBSITE := anatma.github.io

all: build

build: deps test
	go build -ldflags "-X main.version=$(VERSION)" github.com/anatma/autotune/cmd/autotune
	go build -ldflags "-X CmdName=autotune-pro -X main.version=$(VERSION)" github.com/anatma/autotune/cmd/autotune-pro
	$(MAKE) website-assets

deps:
	go get ./...

test:
	golint ./...
	go test -cover ./...
	go tool vet **/*.go

archive:
	tar cvzf $(PRODUCT)-$(VERSION).tar.gz $(PRODUCT)

release: spell build archive
	sed -i -e "s/^VERSION=.*$\/VERSION=$(VERSION)/g" website/install.sh
	git add website/install.sh
	-git commit -m "Version $(VERSION)"
	-git tag v$(VERSION) && git push --tags
	s3cmd put --acl-public $(PRODUCT)-$(VERSION).tar.gz s3://assets.anatma.co/$(PRODUCT)/${VERSION}/$(PRODUCT)-${VERSION}.tar.gz

website-assets:
	jq -n --arg PROFILES "$(PROFILES)" '$$PROFILES | split(" ")' > website/js/profiles.json

website:
	echo "Nothin here govn'r"

website-dev:  website
	rm -rf $$GOPATH/src/github.com/anatma/$(WEBSITE)/content/$(PRODUCT)
	cp -r ./website $$GOPATH/src/github.com/anatma/$(WEBSITE)/content/$(PRODUCT)

spell:
	for i in $(shell ls website/*.html); do \
		aspell check --mode=html $$i; \
	done

.PHONY: website website-dev
