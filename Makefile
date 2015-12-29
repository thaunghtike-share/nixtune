all:
	go get
	golint **/*.go
	golint *.go
	go build
