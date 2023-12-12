FROM golang:1.20-alpine

WORKDIR /workspace

COPY bin/app .

ENTRYPOINT ["./app"]