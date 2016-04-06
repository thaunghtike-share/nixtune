PRODUCT := autotune
VERSION := $(shell cat VERSION)
WEBSITE := acksin.com

all: build

build: deps test
	echo "Lambda API Key: $(AUTOTUNE_LAMBDA_API_KEY)"
	echo "Lambda Secret Key: $(AUTOTUNE_LAMBDA_SECRET_KEY)"
	echo "Lambda Region: $(AUTOTUNE_LAMBDA_REGION)"
	go build -ldflags "-X main.version=$(VERSION) -X main.awsAPIKey=$(AUTOTUNE_LAMBDA_API_KEY) -X main.awsSecretKey=$(AUTOTUNE_LAMBDA_SECRET_KEY) -X main.awsRegion=$(AUTOTUNE_LAMBDA_REGION)"
	$(MAKE) website-assets

deps:
	go get -u github.com/golang/lint/golint
	go get -u github.com/jteeuwen/go-bindata/...
	go-bindata -ignore="signatures/README.md" -o sigfiles.go signatures/open signatures/pro signatures/premium 
	go get ./...

dev-deps:
	sudo apt-get install -y inkscape 

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
	s3cmd put --acl-public $(PRODUCT)-$(VERSION).tar.gz s3://assets.acksin.co/$(PRODUCT)/${VERSION}/$(PRODUCT)-${VERSION}.tar.gz

website-assets:
	# cd website && go run logo.go > logo.svg && inkscape -z -d 150 -e autotune.png logo.svg
	./autotune list > website/signatures.json
	emacs DOCUMENTATION.org --batch --eval '(org-html-export-to-html nil nil nil t)'  --kill
	echo "---" > website/docs.html.erb
	echo "title: Acksin Autotune Docs" >> website/docs.html.erb
	echo "layout: docs" >> website/docs.html.erb
	echo "---" >> website/docs.html.erb
	cat DOCUMENTATION.html >> website/docs.html.erb
	rm DOCUMENTATION.html

website:
	echo "Nothin here govn'r"

website-dev:  website
	rm -rf $$GOPATH/src/github.com/acksin/$(WEBSITE)/content/$(PRODUCT)
	cp -r ./website $$GOPATH/src/github.com/acksin/$(WEBSITE)/content/$(PRODUCT)

spell:
	for i in $(shell ls website/*.html); do \
		aspell check --mode=html $$i; \
	done

.PHONY: website website-dev
