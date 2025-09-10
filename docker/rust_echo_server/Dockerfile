FROM rust:1.89.0 as builder

ARG GIT_DESCRIBE

ENV SOURCE_ROOT /src
WORKDIR ${SOURCE_ROOT}

COPY rust/ ${SOURCE_ROOT}/

RUN rustup target add x86_64-unknown-linux-musl
RUN GIT_DESCRIBE=${GIT_DESCRIBE} cargo build --release --target x86_64-unknown-linux-musl --bin echo_server

FROM alpine:latest

ENV APP_HOME /app
ENV SOURCE_ROOT /src
WORKDIR $APP_HOME

COPY --from=builder ${SOURCE_ROOT}/target/x86_64-unknown-linux-musl/release/echo_server ${APP_HOME}

COPY docker/rust_echo_server/config.json ${APP_HOME}

EXPOSE 80
CMD ["./echo_server", "-c", "config.json"]
