FROM golang:1.20 as builder

ENV APP_HOME /app
WORKDIR "$APP_HOME"

COPY go/cmd/echo_server/ .
COPY go/go.mod .

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 go build -o echo_server

FROM alpine:latest

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY --from=builder "$APP_HOME"/echo_server $APP_HOME

COPY docker/go_echo_server/config.json $APP_HOME

EXPOSE 80
CMD ["./echo_server", "-c", "config.json"]
