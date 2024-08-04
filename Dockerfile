FROM golang:1.22.5-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o management-service ./cmd

FROM alpine as runner

COPY --from=builder /app/management-service .
COPY config/config.yaml ./config/config.yaml

CMD ["./management-service"]