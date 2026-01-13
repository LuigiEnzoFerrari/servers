FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o user-server ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/user-server .

EXPOSE 50051

CMD ["./user-server"]