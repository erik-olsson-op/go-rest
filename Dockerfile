FROM golang:1.23

WORKDIR /app
COPY . .
RUN go mod download && go mod verify
RUN go build -C cmd/ -v -o api

EXPOSE 8080

CMD ["./cmd/api"]