FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go mod tidy
EXPOSE 8080

CMD [ "go", "run", "main.go" ]