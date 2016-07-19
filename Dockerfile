FROM alpine:3.4
RUN apk update && \
    apk add ca-certificates
ADD network-test /test
CMD ["/test"]
