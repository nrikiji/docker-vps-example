FROM golang:1.15.11

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

EXPOSE 80

# HEALTHCHECK CMD curl --fail http://localhost:8080/ || exit 1

CMD ["./main"]
