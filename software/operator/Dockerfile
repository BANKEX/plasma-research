FROM golang:1.10.3 as builder
RUN wget https://www.foundationdb.org/downloads/5.1.7/ubuntu/installers/foundationdb-clients_5.1.7-1_amd64.deb && dpkg -i foundationdb-clients_5.1.7-1_amd64.deb
RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
# apk add --no-cache make gcc musl-dev linux-headers
WORKDIR /go/src/github.com/poma/plasma-backend/
COPY . .
#RUN go get -d -v
RUN dep ensure -v
RUN go build

FROM ubuntu
#RUN apk --no-cache add ca-certificates
RUN apt update && apt install -y curl wget
RUN wget https://www.foundationdb.org/downloads/5.1.7/ubuntu/installers/foundationdb-clients_5.1.7-1_amd64.deb && dpkg -i foundationdb-clients_5.1.7-1_amd64.deb
WORKDIR /root/
COPY --from=builder /go/src/github.com/poma/plasma-backend/plasma-backend .
COPY fdb.cluster /etc/foundationdb/fdb.cluster
CMD ["/root/plasma-backend"]