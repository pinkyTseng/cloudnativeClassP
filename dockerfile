
FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod /app
RUN go mod download

COPY serverpractice/*.go /app

RUN go build -o /app/httpserver


FROM busybox
WORKDIR /app
COPY --from=build /app/httpserver /app/
EXPOSE 8080
ENTRYPOINT [ "./httpserver" ]