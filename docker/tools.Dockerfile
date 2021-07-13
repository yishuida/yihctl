FROM golang:1.16-alpine3.14 as builder

ENV GOPROXY https://goproxy.io
WORKDIR /go/src/daqiang.dev/tools/

COPY . .

RUN CGO_ENABLED=0 go install -ldflags '-s -w' server.go

FROM alpine:3.14

COPY --from=builder /go/bin/server /server

RUN adduser -h /home/client -s /bin/sh -u 1001 -D client && \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories && \
    apk upgrade && \
    apk add --no-cache curl wget

USER 1001
EXPOSE 8080

ENTRYPOINT [ "/server" ]