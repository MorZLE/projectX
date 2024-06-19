# Use the official Ubuntu image
FROM golang:alpine

WORKDIR /app

COPY . .
# Install dependencies
RUN go mod download

# Build the application
RUN go build -o bot msrvs/msrv-bot-tg/cmd/main.go

CMD ["./bot"]