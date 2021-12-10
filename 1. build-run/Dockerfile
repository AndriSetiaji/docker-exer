#!/bin/bash
FROM alpine:latest
RUN apk add --no-cache ca-certificates

RUN apk add --no-cache tzdata

WORKDIR /app

ADD main /app/main

ENTRYPOINT ["/app/main"]
