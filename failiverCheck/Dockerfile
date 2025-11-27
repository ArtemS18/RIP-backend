FROM golang:1.25 AS builder

ENV CGO_ENABLED=0
ENV GOOS=linux

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY templates resources ./
COPY . .
RUN go build -ldflags="-s -w" -o app ./cmd/

FROM alpine
WORKDIR /app

COPY --from=builder /build/app ./app
COPY --from=builder /build/templates ./templates 
COPY --from=builder /build/resources ./resources 
ENTRYPOINT ["./app"]
