FROM scratch
ADD bin/templeton_linux_amd64 /templeton
CMD ["/templeton"]
