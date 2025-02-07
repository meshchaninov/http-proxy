FROM golang:latest as builder
WORKDIR /go/src/github.com/meshchaninov/http-proxy
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-s -w' -o ./proxy

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/meshchaninov/http-proxy/proxy /proxy
ENTRYPOINT ["/proxy"]
