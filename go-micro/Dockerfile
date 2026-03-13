FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./server.go

FROM alpine:latest
WORKDIR /root/
# install CA certs for JWT/HTTPS
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/server .

EXPOSE 8080
CMD ["./server"]