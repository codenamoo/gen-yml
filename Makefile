.PHONY: run build get test install clean

GO=${shell which go}

VERSION_MAJOR=0
VERSION_MINOR=0
VERSION_COUNT=1

PREFIX?=/usr/local/gen-yml/
BIN_PATH=$(PREFIX)/bin
INSTALL?= cp -a

EXE_NAME=gen-yml
VERSION_NAME=$(EXE_NAME)-$(VERSION_MAJOR).$(VERSION_MINOR).$(VERSION_COUNT)


run: 
	go run main.go

build:
	godep restore
	go build -o bin/$(VERSION_NAME) main.go
	ln -sf $(VERSION_NAME) bin/$(EXE_NAME)

get: 
	go get gopkg.in/yaml.v2
	go get github.com/codegangsta/cli
	go get github.com/codenamoo/typeconv

test:
#	go test -v src/test/point_test.go


install: build
	mkdir -p $(BIN_PATH)
	$(INSTALL) $(VERSION_NAME) $(BIN_PATH)
#	ln -sf $(BIN_PATH)/$(VERSION_NAME) $(BIN_PATH)/$(EXE_NAME)

clean:
	go clean
	rm -rf bin/$(VERSION_NAME) bin/$(EXE_NAME)
