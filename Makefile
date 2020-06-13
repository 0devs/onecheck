build-full:
	go build -o dist/onecheck-full ./cmd/onecheck-full

build-full-release:
	go build -o dist/onecheck-full -ldflags '-s -w' ./cmd/onecheck-full

release: build-full-release

build-plugin-file:
	go build -buildmode=plugin -o dist/file.so ./plugin/file/main