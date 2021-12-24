FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM alpine:3.13

COPY --from=builder /src/bin /app

WORKDIR /app

EXPOSE 31233
EXPOSE 31234

CMD ["./core-broker"]
