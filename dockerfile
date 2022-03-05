
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY serverpractice/*.go ./

RUN go build -o /httpserver

EXPOSE 8080

CMD [ "/docker-gs-ping" ]