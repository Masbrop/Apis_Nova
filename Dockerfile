# syntax=docker/dockerfile:1

FROM golang:1.24-alpine AS builder

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM alpine:3.21

RUN apk add --no-cache ca-certificates && \
    addgroup -S app && \
    adduser -S app -G app

WORKDIR /app

COPY --from=builder /out/api /app/api

USER app

EXPOSE 8080

ENTRYPOINT ["/app/api"]
