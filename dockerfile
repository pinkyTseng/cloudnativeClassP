
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod /app
RUN go mod download

COPY serverpractice/*.go /app

RUN go build -o /httpserver

EXPOSE 8080

ENTRYPOINT [ "/httpserver" ]