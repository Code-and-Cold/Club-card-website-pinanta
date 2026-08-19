FROM golang:1.26.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/main /main
COPY --from=builder /app/static /static

ENTRYPOINT ["/main"]
