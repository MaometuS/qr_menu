FROM golang:1.22-alpine AS builder
WORKDIR /build
COPY . .
RUN go build -o main

FROM alpine
#RUN adduser -S -D -H -h /app appuser
#USER appuser
COPY --from=builder /build/main /app/
COPY .env_piarmenu /app/
WORKDIR /app
CMD ["./main"]
