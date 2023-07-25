.PHONY: build clean

ARCHS := amd64 arm64

build:
	echo "Building for local";
	go build -o bin/ffa_server ./cmd/...

	echo "Building for amd64 linux";
	GOOS=linux GOARCH=amd64 go build -o bin/ffa_server_linux_amd64 ./cmd/...

	echo "Building for arm64 linux";
	GOOS=linux GOARCH=arm64 go build -o bin/ffa_server_linux_arm64 ./cmd/...

clean:
	rm -rf bin/*