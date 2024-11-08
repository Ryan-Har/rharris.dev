FROM golang:1.23 AS builder
WORKDIR /app
COPY src/go.mod src/go.sum ./
RUN go mod download
COPY src/. .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/web/main.go

# Stage 2: Create the minimal final image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY src/static ./static
RUN chmod +x ./main
CMD ["./main"]