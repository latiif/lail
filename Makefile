BINARY = lail
VERIFY_REPORT = vet.report
TEST_REPORT = tests.xml
GOARCH = amd64

VERSION?=v0.1.0
COMMIT=$(shell git rev-parse HEAD)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)

# Symlink into GOPATH
GITHUB_USERNAME=latiif
BUILD_DIR=$(shell pwd)
BIN_DIR=${BUILD_DIR}/bin
CURRENT_DIR=$(shell pwd)
BUILD_DIR_LINK=$(shell readlink ${BUILD_DIR})

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-X github.com/latiif/awaitrmq/cmd.VERSION=${VERSION} -X github.com/latiif/awaitrmq/cmd.COMMIT=${COMMIT} -X github.com/latiif/awaitrmq/cmd.BRANCH=${BRANCH}"

all: clean test vet linux darwin windows

linux:
	cd ${BUILD_DIR}; \
	CGO_ENABLED=0  GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY}-linux-${GOARCH} . ; \
	cd - >/dev/null

darwin:
	cd ${BUILD_DIR}; \
	CGO_ENABLED=0  GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY}-darwin-${GOARCH} . ; \
	cd - >/dev/null

windows:
	cd ${BUILD_DIR}; \
	CGO_ENABLED=0  GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ${BIN_DIR}/${BINARY}-windows-${GOARCH}.exe . ; \
	cd - >/dev/null

wasm:
	cd ${BUILD_DIR}; \
	docker run --rm -v ${CURRENT_DIR}:/src/github.com/latiif/lail tinygo/tinygo:0.13.1 /bin/sh -c "cd /src/github.com/latiif/lail && GOPATH='/' tinygo build -o ${BINARY}-wasm.wasm -target=wasm ." ;\
	cd - >/dev/null


vet:
	-cd ${BUILD_DIR}; \
	go mod verify > ${VERIFY_REPORT} 2>&1 ; \
	cd - >/dev/null

fmt:
	cd ${BUILD_DIR}; \
	go fmt $$(go list ./... | grep -v /vendor/) ; \
	cd - >/dev/null

clean:
	-rm -f ${TEST_REPORT}
	-rm -f ${VERIFY_REPORT}
	-rm -f ${BIN_DIR}/${BINARY}-*
	-rmdir ${BIN_DIR}

.PHONY: linux darwin windows test vet fmt clean
