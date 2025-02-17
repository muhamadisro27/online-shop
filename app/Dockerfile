FROM golang:1.21-alpine as builder

WORKDIR /app

COPY main.go .

RUN go build -o /app/main /app/main.go


FROM alpine:3

WORKDIR /app

COPY --from=builder /app/main /app

CMD /app/main