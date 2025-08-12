FROM golang:1.24.5-alpine AS builder

RUN apk add --no-cache git ca-certificates tzdata
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/ordersystem

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata netcat-openbsd
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app
COPY --from=builder /app/app .
COPY entrypoint.sh .

RUN chmod +x entrypoint.sh
RUN chown -R appuser:appgroup /app

USER appuser
EXPOSE 8080 50051 8081

ENTRYPOINT ["./entrypoint.sh"]
