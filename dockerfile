#FROM golang:1.20.3-alpine AS builder
FROM golang:1.21.4-alpine AS builder

COPY . /github.com/KozlovNikolai/auth/source/
WORKDIR /github.com/KozlovNikolai/auth/source/

RUN go mod download
RUN go build -o ./bin/auth cmd/grpc-server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/KozlovNikolai/auth/source/bin/auth .

CMD ["./auth"]