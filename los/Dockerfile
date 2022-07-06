FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /los

FROM alpine:3.16

WORKDIR /

COPY --from=builder /los /los

EXPOSE 8080

ENTRYPOINT ["/los"]