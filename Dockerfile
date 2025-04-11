FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

COPY .env .env

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/app .

COPY --from=builder /app/.env .env

COPY --from=builder /app/internal/web/templates /root/internal/web/templates

EXPOSE 3001

CMD ["./app"]
