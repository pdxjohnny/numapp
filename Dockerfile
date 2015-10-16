FROM busybox
ADD ./numapp_linux-amd64 /app
CMD ["/app"]
