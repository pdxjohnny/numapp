FROM alpine:latest
MAINTAINER John Andersen <johnandersenpdx@gmail.com>

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY ./static /app/static
COPY ./numapp /app/run

ENTRYPOINT ["/app/run"]
