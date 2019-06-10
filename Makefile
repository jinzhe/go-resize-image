all: build-windows build-linux build-macos
build:
		go build -ldflags "-s -w" main.go

clean:
		go clean
		rm -f resize-windows.exe
		rm -f resize-linux
		rm -f resize-macos
run:
		go run main.go

deps:
		go get github.com/nfnt/resize


# Cross compilation
build-windows:
		GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o resize-windows.exe main.go

build-linux:
		GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o resize-linux main.go

build-macos:
		GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o resize-macos main.go