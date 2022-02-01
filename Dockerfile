# Stage 1 - Build
FROM golang:1.17.4 as builder

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app/ ./cmd/web


# Stage 2 - Run
FROM alpine:3.15

WORKDIR /root/
COPY --from=builder /src .
CMD ["./app/web"]
