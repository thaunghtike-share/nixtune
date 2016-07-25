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

release: website-assets spell build archive
	-git commit -m "Version $(VERSION)"
	-git tag v$(VERSION) && git push --tags
	s3cmd put --acl-public $(PRODUCT)-$(VERSION).tar.gz s3://assets.acksin.com/$(PRODUCT)/${VERSION}/$(PRODUCT)-$(shell uname)-$(shell uname -i)-${VERSION}.tar.gz

website-assets:
	emacs README.org --batch --eval '(org-html-export-to-html nil nil nil t)'  --kill
	echo "---" > docs.html.erb
	echo "title: STRUM Docs" >> docs.html.erb
	echo "layout: docs" >> docs.html.erb
	echo "description: Acksin STRUM documentation for tool that diagnoses Linux issues quickly giving you a complete picture encompassing the CPU, Memory, IO, Networking, Processes, Limits, etc." >> docs.html.erb
	echo "---" >> docs.html.erb
	cat README.html >> docs.html.erb
	rm README.html
	cp docs.html.erb $$GOPATH/src/github.com/acksin/fugue/acksin.com/source/strum/


spell:
	# for i in README.org website/index.html.erb website/_download.erb; do \
	# 	aspell check --dont-backup --mode=html $$i; \
	# done

