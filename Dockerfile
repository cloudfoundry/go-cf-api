# Intermidiate image to build binary
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git make sed curl wget
WORKDIR /tmp/cloudgontroller
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/cloudgontroller ./cmd/main.go

# Final image
FROM scratch
COPY --from=builder /go/bin/cloudgontroller /go/bin/cloudgontroller
CMD ["/go/bin/cloudgontroller"]
