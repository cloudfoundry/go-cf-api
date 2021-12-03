# Intermidiate image to build binary
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git make sed curl wget
WORKDIR /tmp/go-cf-api
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/go-cf-api ./cmd/main.go

# Final image
FROM scratch
COPY --from=builder /go/bin/go-cf-api /go/bin/go-cf-api
CMD ["/go/bin/go-cf-api"]
