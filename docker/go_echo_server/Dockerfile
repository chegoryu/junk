FROM golang:1.20 as builder

ARG GIT_DESCRIBE
ARG PKG

ENV LDFLAGS="-X ${PKG}/pkg/buildinfo.GitDescribe=${GIT_DESCRIBE}"

ENV BUILD_ROOT /build
ENV SOURCE_ROOT /src
WORKDIR ${SOURCE_ROOT}/cmd/echo_server

COPY go/ ${SOURCE_ROOT}/

RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 go build -ldflags "${LDFLAGS}" -o $BUILD_ROOT/echo_server

FROM alpine:latest

ENV APP_HOME /app
ENV BUILD_ROOT /build
WORKDIR $APP_HOME

COPY --from=builder ${BUILD_ROOT}/echo_server ${APP_HOME}

COPY docker/go_echo_server/config.json ${APP_HOME}

EXPOSE 80
CMD ["./echo_server", "-c", "config.json"]
