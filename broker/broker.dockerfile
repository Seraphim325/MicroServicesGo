FROM golang:1.26.1-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o broker ./cmd/api

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/broker /app/

COPY .env /app/

WORKDIR /app

CMD [ "./broker" ]