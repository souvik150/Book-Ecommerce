FROM golang:alpine

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go build -o main ./src/cmd/main.go

CMD ["/app/main"]