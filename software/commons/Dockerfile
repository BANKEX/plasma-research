FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app/operator

RUN go get github.com/gin-gonic/gin
RUN go get github.com/ethereum/go-ethereum
RUN go get github.com/syndtr/goleveldb/leveldb
RUN go get github.com/swaggo/swag
RUN go get github.com/swaggo/gin-swagger/swaggerFiles
RUN go get github.com/swaggo/gin-swagger
RUN go get github.com/gin-gonic/contrib/cors
RUN go get github.com/c-bata/go-prompt

#RUN go build operator.go

EXPOSE 8080

CMD ["go", "run", "operator.go"]
