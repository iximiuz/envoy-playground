FROM golang:1.15

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY main.go .

ENV FAIL_PCT=30

ENV PORT_API=80
EXPOSE 80

ENV PORT_METRICS=8081
EXPOSE 8081

RUN go get ./...

CMD ["go", "run", "main.go"]

