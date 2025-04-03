FROM golang:1.24.1-alpine AS builder

COPY . /github.com/vrnxx/auth/source/
WORKDIR /github.com/vrnxx/auth/source/

RUN go mod download
RUN go build -o ./bin/crud_server cmd/grpc/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/vrnxx/auth/source/bin/crud_server .
COPY --from=builder /github.com/vrnxx/auth/source/internal/configs/app/ .
RUN export CONFIG_PATH=""
CMD ["./crud_server"]