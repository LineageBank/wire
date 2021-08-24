FROM golang:1.17-alpine as builder
WORKDIR /go/src/github.com/moov-io/wire
RUN apk add -U --no-cache make bash gcc git
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
RUN make build
USER moov

FROM scratch
LABEL maintainer="Moov <support@moov.io>"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/wire/bin/server /bin/server
COPY --from=builder /etc/passwd /etc/passwd

USER moov
EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
