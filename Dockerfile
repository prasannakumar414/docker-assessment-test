FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/server .

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /out/server /app/server

EXPOSE 8080

ENTRYPOINT ["/app/server"]
