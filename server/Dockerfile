FROM golang:1.21
WORKDIR /app
COPY main.go .
RUN go build -o server main.go
EXPOSE 8081
CMD ["./server"]