FROM golang:1.25.3-alpine AS builder

WORKDIR /app

# Install dependensi tambahan (git & tzdata)
RUN apk add --no-cache git tzdata

# Copy go.mod & go.sum dulu untuk caching build
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN go build -o booking-fields-app  .

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache postgresql-client bash
COPY --from=builder /app/booking-fields-app .
COPY wait-for-postgresql.sh .
RUN chmod +x wait-for-postgresql.sh

# Jalankan aplikasi
CMD ["./wait-for-postgresql.sh", "db", "./booking-fields-app"]