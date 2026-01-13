FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o dashboard-server ./cmd/api/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/dashboard-server .

EXPOSE 50051

CMD ["./dashboard-server"]