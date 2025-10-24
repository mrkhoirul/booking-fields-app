FROM golang:1.25.3-alpine

WORKDIR /app

# Install dependensi tambahan (git & tzdata)
RUN apk add --no-cache git tzdata

# Copy go.mod & go.sum dulu untuk caching build
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN go build -o /booking-fields-app

EXPOSE 8080

# Jalankan aplikasi
CMD ["/booking-fields-app"]