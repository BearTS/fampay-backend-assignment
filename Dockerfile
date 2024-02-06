# Stage 1: Build the application
FROM golang:1.21-alpine AS builder

RUN mkdir /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine
RUN apk add tzdata
ENV TZ=Asia/Kolkata

WORKDIR /app
COPY --from=builder /app/main .
USER nobody:nobody
ENTRYPOINT ["./main"]