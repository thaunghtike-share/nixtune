PRODUCT := strum
VERSION := $(shell cat VERSION)
WEBSITE := acksin.com

all: build

build: deps test
	go build -ldflags "-X main.version=$(VERSION)"
	$(MAKE) website-assets

deps:
	go get -u github.com/golang/lint/golint
	go get ./...

dev-deps:
	# sudo apt-get install -y inkscape 
	sudo add-apt-repository -y ppa:ubuntu-elisp/ppa && sudo apt-get -qq update && sudo apt-get -qq -f install && sudo apt-get -qq install emacs-snapshot && sudo apt-get -qq install emacs-snapshot-el;
	emacs --version
	wget https://raw.githubusercontent.com/acksin/release-checklist/master/install-orgmode.el
	emacs-snapshot --batch -l ./install-orgmode.el

test:
	golint ./...
	go test -cover ./...
	go tool vet **/*.go

archive:
	tar cvzf $(PRODUCT)-$(VERSION).tar.gz $(PRODUCT)

release: spell build archive
	sed -i -e "s/^VERSION=.*$\/VERSION=$(VERSION)/g" website/install.sh
	sed -i -e "s/^version: .*$\/version: $(VERSION)/g" website/index.html
	git add website/install.sh
	-git commit -m "Version $(VERSION)"
	-git tag v$(VERSION) && git push --tags
	s3cmd put --acl-public $(PRODUCT)-$(VERSION).tar.gz s3://assets.acksin.com/$(PRODUCT)/${VERSION}/$(PRODUCT)-$(shell uname)-$(shell uname -i)-${VERSION}.tar.gz

website-assets:
	emacs DOCUMENTATION.org --batch --eval '(org-html-export-to-html nil nil nil t)'  --kill
	echo "---" > website/docs.html.erb
	echo "title: Acksin STRUM Docs" >> website/docs.html.erb
	echo "layout: docs" >> website/docs.html.erb
	echo "description: Acksin STRUM documentation for tool that diagnoses Linux issues quickly giving you a complete picture encompassing the CPU, Memory, IO, Networking, Processes, Limits, etc." >> website/docs.html.erb
	echo "---" >> website/docs.html.erb
	cat DOCUMENTATION.html >> website/docs.html.erb
	rm DOCUMENTATION.html

website:
	echo "Nothin here govn'r"

spell:
	for i in $(shell ls website/*.html); do \
		aspell check --mode=html $$i; \
	done

.PHONY: website website-dev
