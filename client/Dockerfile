FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go .
RUN go build -o client main.go
EXPOSE 8080
CMD ["./client"]