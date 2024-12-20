FROM golang:latest

WORKDIR /app
COPY . .
RUN go install github.com/air-verse/air@latest
RUN go build

CMD ["/go/bin/air"]
# CMD ["tail", "-f"]