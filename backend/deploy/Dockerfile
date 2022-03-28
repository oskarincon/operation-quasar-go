FROM golang:1.17

WORKDIR /api
COPY . /api

RUN go build -o main ./main.go
CMD ["./main"]