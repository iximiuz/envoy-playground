FROM golang:1.15

ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY main.go .

ENV PORT_API=80
EXPOSE 80

ENV PORT_METRICS=8088
EXPOSE 8088

RUN go get ./...

CMD ["go", "run", "main.go"]

