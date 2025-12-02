FROM golang:1.25-trixie AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM debian:trixie-slim

RUN apt-get update && apt-get install -y ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/templates ./templates

EXPOSE 8081

CMD ["./main"]
