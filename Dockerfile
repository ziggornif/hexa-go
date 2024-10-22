FROM golang:1-alpine as builder
RUN apk update && apk add upx gcc make g++

WORKDIR /build
ADD . .
RUN make build

FROM alpine
COPY --from=builder /build/hexa-go /bin/hexa-go
RUN chmod +x /bin/hexa-go

ENV GIN_MODE=release

ENTRYPOINT ["/bin/hexa-go"]
