all:
	go get
	golint **/*.go
	golint *.go
	go test -cover
	go vet
	go build
