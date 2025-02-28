FROM golang:1.16.3-alpine3.13 AS builder

WORKDIR /app

COPY main.go .
RUN go build main.go

FROM golang:1.16.3-alpine3.13 AS runner

WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]