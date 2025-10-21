FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main .
FROM alpine:latest
WORKDIR /
COPY --from=builder /main /main
EXPOSE 8080
CMD ["/main"]