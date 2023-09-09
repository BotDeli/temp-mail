FROM golang:latest
WORKDIR /app
COPY . .
EXPOSE 3000
CMD go run cmd/main/main.go