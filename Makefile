# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
BINARY_NAME=fanyi

all: build-all

# 清理目标：清理生成的二进制文件
clean:
	rm -f bin/$(BINARY_NAME)-windows-amd64.exe
	rm -f bin/$(BINARY_NAME)-linux-amd64
	rm -f bin/$(BINARY_NAME)-darwin-amd64

# 格式化目标：格式化源代码文件
fmt:
	$(GOFMT) .\...

# 交叉编译目标：构建适用于 Linux 的二进制文件
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-linux-amd64 -v .
	CGO_ENABLED=0 GOOS=linux GOARCH=386 $(GOBUILD) -o bin/$(BINARY_NAME)-linux-386 -v .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm $(GOBUILD) -o bin/$(BINARY_NAME)-linux-arm -v .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 $(GOBUILD) -o bin/$(BINARY_NAME)-linux-arm64 -v .
# 交叉编译目标：构建适用于 Windows 的二进制文件
build-windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-windows-amd64.exe -v .
	CGO_ENABLED=0 GOOS=windows GOARCH=386 $(GOBUILD) -o bin/$(BINARY_NAME)-windows-386.exe -v .
	CGO_ENABLED=0 GOOS=windows GOARCH=arm $(GOBUILD) -o bin/$(BINARY_NAME)-windows-arm.exe -v .
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 $(GOBUILD) -o bin/$(BINARY_NAME)-windows-arm64.exe -v .
# 交叉编译目标：构建适用于 macOS 的二进制文件
build-mac:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-darwin-amd64 -v .
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 $(GOBUILD) -o bin/$(BINARY_NAME)-darwin-arm64 -v .

# 交叉编译目标：构建适用于 FreeBSD 的二进制文件

build-freebsd:
	CGO_ENABLED=0 GOOS=freebsd GOARCH=amd64 $(GOBUILD) -o bin/$(BINARY_NAME)-freebsd-amd64 -v .
	CGO_ENABLED=0 GOOS=freebsd GOARCH=386 $(GOBUILD) -o bin/$(BINARY_NAME)-freebsd-386 -v .
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm $(GOBUILD) -o bin/$(BINARY_NAME)-freebsd-arm -v .
	CGO_ENABLED=0 GOOS=freebsd GOARCH=arm64 $(GOBUILD) -o bin/$(BINARY_NAME)-freebsd-arm64 -v .
# 构建所有平台的程序
build-all: build-linux build-windows build-mac build-freebsd
	@echo "Compiling for every OS and Platform"

# 为了确保这些目标不会与文件名冲突，将它们声明为伪目标
.PHONY: all build test clean fmt build-linux build-windows build-mac build-all