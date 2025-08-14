# Build stage
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o publisher .

# Runtime stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/publisher .
CMD ["./publisher", "Mensaje desde Docker!"]