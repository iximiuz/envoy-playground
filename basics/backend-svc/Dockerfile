FROM golang:1.15

ENV ERROR_RATE=30

ENV PORT_API=80
EXPOSE 80

ENV PORT_METRICS=8081
EXPOSE 8081

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go mod download
RUN go build -o backend

CMD ["/app/backend"]

