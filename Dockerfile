FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /api ./...

FROM alpine:latest

COPY --from=builder /api /api

CMD /api
