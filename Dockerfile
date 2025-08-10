# syntax=docker/dockerfile:1.2
FROM alpine:3.22

RUN apk add --no-cache --no-progress ca-certificates tzdata

ARG TARGETPLATFORM
COPY ./dist/$TARGETPLATFORM/apache4 /

EXPOSE 80
VOLUME ["/tmp"]

ENTRYPOINT ["/apache4"]
