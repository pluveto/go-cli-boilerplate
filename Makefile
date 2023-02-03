BIN_NAMES=greet
GOARCHS=amd64 386 arm arm64

dev: linux

default: all

all: windows linux mac

prepare:
	@mkdir -p bin

windows: prepare
	GOOS=windows
	
	for BIN_NAME in $(BIN_NAMES); do \
		for GOARCH in $(GOARCHS); do \
			GOOS=windows GOARCH=$$GOARCH go build -o bin/$$GOOS_$$GOARCH/$$BIN_NAME.exe cmd/$$BIN_NAME/main.go; \
		done \
	done

linux: prepare
	GOOS=linux
	for BIN_NAME in $(BIN_NAMES); do \
		for GOARCH in $(GOARCHS); do \
			GOOS=linux GOARCH=$$GOARCH go build -o bin/$$GOOS_$$GOARCH/$$BIN_NAME cmd/$$BIN_NAME/main.go; \
		done \
	done

mac: prepare
	GOOS=darwin
	for BIN_NAME in $(BIN_NAMES); do \
		for GOARCH in $(GOARCHS); do \
			GOOS=darwin GOARCH=$$GOARCH go build -o bin/$$GOOS_$$GOARCH/$$BIN_NAME cmd/$$BIN_NAME/main.go; \
		done \
	done

.PHONY: all, default