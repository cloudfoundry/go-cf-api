# Intermidiate image to build binary
FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git make sed curl wget
WORKDIR /tmp/gopilot
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/gopilot ./cmd/main.go

# Final image
FROM scratch
COPY --from=builder /go/bin/gopilot /go/bin/gopilot
CMD ["/go/bin/gopilot"]
