FROM busybox
ADD ./numapp_linux-amd64 /app/run
ADD ./keys/http /app/keys/http
ADD ./static /app/static
CMD ["/app/run"]
