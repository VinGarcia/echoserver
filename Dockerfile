FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /api ./...

FROM alpine:latest

ARG USER=nonroot

# Create the non root user
RUN adduser -D $USER -u 1000

COPY --from=builder /api /api

USER 1000

CMD /api
