FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app

FROM ubuntu:22.04

WORKDIR /root/
COPY --from=builder /app/app .
COPY --from=builder /app/gogogadget.yaml .
COPY --from=builder /app/views/ ./views/
 
EXPOSE 8081

CMD ["./app"]


