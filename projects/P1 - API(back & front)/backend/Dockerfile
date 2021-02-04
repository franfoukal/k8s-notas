FROM golang:alpine as builder
# Multi stage building
WORKDIR /app
COPY ./src/main.go .
# Compila el código y genera el ejecutable 'app'
RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
# Corre el ejecutable generado
CMD ["./app"]