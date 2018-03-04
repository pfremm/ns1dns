FROM golang:1.9.2-alpine

RUN apk add --no-cache --update alpine-sdk

COPY . /go/src/github.com/pfremm/ns1dns
WORKDIR /go/src/github.com/pfremm/ns1dns
RUN apk add --update ca-certificates openssl \
    && make release-binary \
    && cp /go/src/github.com/pfremm/ns1dns/bin/ns1dns /usr/local/bin/ns1dns

WORKDIR /

ENTRYPOINT ["ns1dns"]

CMD ["version"]