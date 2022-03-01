FROM alpine:3.9
RUN apk --no-cache add ca-certificates
RUN apk --no-cache add tzdata
COPY http-echo /
ENTRYPOINT ["/http-echo"]