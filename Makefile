build-full:
	go build -o dist/onecheck-full ./cmd/onecheck-full

build-plugin-file:
	go build -buildmode=plugin -o dist/file.so ./plugin/file/main