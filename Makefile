.PHONY: clean

prepare:
	mkdir -p bin/

build: prepare
	GOOS=linux GOARCH=amd64 GOPROXY="https://goproxy.cn" go build  -o bin/yihctl-linux

clean:
	rm -rf bin