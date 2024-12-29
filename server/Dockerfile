FROM golang:latest

COPY ./srcs /app
WORKDIR /app
RUN go install github.com/air-verse/air@latest
ENV GIN_MODE=release
# RUN go mod tidy
# RUN go build

CMD ["/go/bin/air"]
