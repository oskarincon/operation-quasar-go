FROM golang:1.17

WORKDIR /app
COPY . /app

RUN go build -o main ./main.go
CMD ["./main"]