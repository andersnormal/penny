FROM scratch
ADD bin/templeton_*_linux_amd64 /templeton
CMD ["/templeton"]
