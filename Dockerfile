FROM golang:1.16 AS builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o vaccine .

FROM alpine:3.13

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" >  /etc/timezone

WORKDIR /usr/src/app

COPY --from=builder /src/vaccine /usr/src/app/vaccine

RUN apk add dumb-init
ENTRYPOINT ["/usr/bin/dumb-init", "--"]

EXPOSE 8080
CMD ["./vaccine"]