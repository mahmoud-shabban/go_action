FROM golang:tip-alpine3.22 AS builder

RUN mkdir /app

WORKDIR /app

COPY go.mod  /app

COPY main.go /app

RUN go build -o greet .

FROM  gcr.io/distroless/base-debian12

WORKDIR /

COPY --from=builder /app/greet .

ENTRYPOINT ["./greet"]
