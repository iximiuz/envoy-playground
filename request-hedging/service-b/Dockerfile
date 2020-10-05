FROM golang:1.15

# Build
ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY main.go .

RUN go mod download
RUN go build -o service-b

# Run
EXPOSE 80
EXPOSE 8081

ENV SERVICE_PORT=80
ENV METRICS_PORT=8081

CMD ["/app/service-b"]

