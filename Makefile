all: clean vendor build
.PHONY: init vendor build run clean lint
init:
	go mod init github.com/futuretea/yapidoc
vendor:
	GOPROXY=https://goproxy.io go mod vendor
docker:
	docker build -t futuretea/yapidoc .
build:
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo -o yapidoc
run:
	go run main.go
test:
	docker run -it --rm --net=host futuretea/yapidoc
lint:
	golangci-lint run
clean:
	rm -rf yapidoc
