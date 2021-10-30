BINARY="./bin/api-server"

setup:
	mkdir -p bin

build: setup
	go build -o ${BINARY} ./cmd/server

run: build
	${BINARY}
