FROM golang:1.24.0 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

# Use a minimal base image for the final container
FROM gcr.io/distroless/base-debian12

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
