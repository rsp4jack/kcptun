FROM golang:1.20.6-alpine3.18 as builder
MAINTAINER xtaci <daniel820313@gmail.com>
ENV GO111MODULE=on
ENV VERSION=`date -u +%Y%m%d`
RUN apk update && \
    apk upgrade && \
    apk add git gcc libc-dev linux-headers
RUN mkdir /build
RUN cd /build
RUN git clone https://github.com/xtaci/kcptun
RUN cd kcptun
RUN go build -mod=vendor -ldflags "-X main.VERSION=$VERSION -s -w" -o /build/client github.com/xtaci/kcptun/client
RUN go build -mod=vendor -ldflags "-X main.VERSION=$VERSION -s -w" -o /build/server github.com/xtaci/kcptun/server

FROM alpine:3.18
RUN apk add --no-cache iptables
COPY --from=builder /build/client /bin
COPY --from=builder /build/server /bin
EXPOSE 29900/udp
EXPOSE 12948
