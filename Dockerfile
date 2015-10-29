FROM busybox
WORKDIR /app
ADD ./keys /app/keys
RUN mkdir -pv /etc/ssl/ && \
    ln -s  /app/keys/etc/ssl/certs /etc/ssl/certs
ADD ./numapp_linux-amd64 /app/run
ADD ./static /app/static
CMD ["/app/run", "http"]
