.PHONY: clean build linux windows test

clean:
	rm -rf ./bin/steg*

build:
	make clean
	mkdir -p bin/
	CGO_ENABLED=0 \
	go build -o ./bin/steg main/main.go

linux:
	make clean
	mkdir -p bin/
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./bin/steg_linux main/main.go

windows:
	make clean
	mkdir -p bin/
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 \
	go build -o ./bin/steg_win.exe main/main.go

test:
	go test $(shell find tests -name "*_test.go")
