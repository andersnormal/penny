FROM scratch
ADD bin/penny_linux_amd64 /penny
CMD ["/penny"]
