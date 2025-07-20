FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN go mod init example.com/myapp && go get -u github.com/gorilla/mux

COPY main.go .

RUN go build -o main main.go

FROM alpine:3.20 AS base

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]