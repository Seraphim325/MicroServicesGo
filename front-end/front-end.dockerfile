FROM golang:1.26.1-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o frontend ./cmd/web

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/frontend /app/

COPY --from=builder /app/cmd/web/templates /app/templates/

COPY .env /app/

WORKDIR /app

CMD [ "./frontend" ]
