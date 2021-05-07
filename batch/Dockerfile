FROM golang:1.15.11 as build-stage

WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

FROM alpine:latest as production-stage
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
COPY cron /var/spool/cron/crontabs/root

COPY --from=build-stage /go/src/app /usr/local/bin
RUN chmod 0755 /usr/local/bin/main
CMD ["crond","-f", "-L", "/dev/stdout"]
