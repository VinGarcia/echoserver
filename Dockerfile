FROM golang:1.20-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o /api ./...

FROM alpine:latest

ARG USER=nonroot

# Create the non root user
RUN adduser -D $USER

COPY --from=builder /api /api

USER $USER

CMD /api
