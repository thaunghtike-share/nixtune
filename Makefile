all:
	go get
	golint ./...
	go test -cover
	go vet
	go build
