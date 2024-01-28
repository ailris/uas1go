FROM golang:1.16-alpine

WORKDIR /app

COPY . .

ENV GO111MODULE=off

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

EXPOSE 8080

CMD ["/app/main"]
