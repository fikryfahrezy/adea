FROM golang:1.18.3-alpine3.16 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /los

FROM alpine:3.16

WORKDIR /app

COPY --from=builder /los /app/los

COPY ./tmp/ ./tmp

EXPOSE 4000

ENTRYPOINT ["/app/los"]