GO_CMD=go
GO_OS=linux
GO_ARCH=amd64
GO_BUILD=GOOS=$(GO_OS) GOARCH=$(GO_ARCH) $(GO_CMD) build

OUT = cgo-sample

SRC = main.go

all:
	$(GO_BUILD) -o $(OUT)

clean:
	$(GO_CMD) clean
	rm -f $(OUT)
