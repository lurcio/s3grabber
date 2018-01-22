
VERSION=$(shell git fetch --tags && git describe --tags --candidates=1 --dirty --always)
FLAGS=-s -w -X main.Version=$(VERSION)
BIN=build/s3grabber-$(shell uname -s)-$(shell uname -m)-$(VERSION)
SRC=$(shell find . -name '*.go')

build: $(BIN)

clean:
	-rm -rf build/

$(BIN): $(SRC)
	-mkdir -p build/
	go build -o $(BIN) -ldflags="$(FLAGS)" .
ifeq ($(shell command -v upx),)
	@echo "No upx in PATH, consider installing it to generate a compressed binary"
else
	upx -9 $(BIN)
endif
