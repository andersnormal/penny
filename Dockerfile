FROM alpine:3.7
ADD bin/templeton_linux_amd64 /templeton
CMD ["/templeton"]
