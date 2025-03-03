FROM golang:1.23-bookworm

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o app

CMD ["/build/app"]