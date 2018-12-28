FROM golang:1.11-alpine

RUN apk add --update --no-cache make linux-headers libc-dev gcc git gmp gmp-dev

WORKDIR /app/
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download

ADD . /app/
RUN make build/operator

EXPOSE 3001
EXPOSE 8545

CMD ["/app/target/operator_linux_x86_64"]
