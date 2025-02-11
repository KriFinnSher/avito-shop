FROM golang:1.23

WORKDIR /avito-shop
COPY . .

RUN go mod tidy
RUN go build -o /build ./cmd/server

EXPOSE 8080
CMD ["/build"]
