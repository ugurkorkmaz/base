# Go Builder Stage
FROM golang:1.20 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o backend .

# Final Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app/

COPY --from=builder /app/backend .

EXPOSE 8080
ENTRYPOINT [ "./backend" ]