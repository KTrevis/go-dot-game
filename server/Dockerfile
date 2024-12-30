FROM golang:latest

COPY ./srcs /app
WORKDIR /app
ENV GIN_MODE=release
# RUN go mod tidy
# RUN go build

CMD ["./air"]
