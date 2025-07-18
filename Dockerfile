FROM golang:1.24.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /app/marketplace ./cmd/main.go

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/marketplace .
COPY --from=builder /app/.env .  

RUN apk add --no-cache tzdata ca-certificates

RUN adduser -D -g '' appuser && chown -R appuser /app
USER appuser

EXPOSE 8080
CMD ["./marketplace"]