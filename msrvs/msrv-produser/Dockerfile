# Use the official Ubuntu image
FROM golang:alpine

WORKDIR /app

COPY . .
# Install dependencies
RUN go mod download

expose 8080
# Build the application
RUN go build -o server msrvs/msrv-produser/cmd/main.go

CMD ["./server"]