FROM golang:1.11-alpine as builder

RUN apk update && apk add git build-base ca-certificates
RUN adduser -D -g 'quotes' quotes
WORKDIR /usr/local/go/src/github.com/bbriggs/quotes
COPY . .
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -ldflags "-s -w -extldflags '-static'" -o /go/bin/quotes

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/quotes /go/bin/quotes
USER quotes
ENTRYPOINT ["/go/bin/quotes"]

