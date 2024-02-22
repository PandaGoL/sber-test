FROM golang:1.21.0-alpine
ENV GOPATH=/
COPY ./ ./

RUN go mod download
RUN go build -o app ./cmd/server/main.go

EXPOSE 8000

ENTRYPOINT ["./app", "-config", "api-sber-test"]
