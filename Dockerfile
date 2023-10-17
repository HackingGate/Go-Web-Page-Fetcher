FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o fetch .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/fetch /app/

ENTRYPOINT ["./fetch"]
