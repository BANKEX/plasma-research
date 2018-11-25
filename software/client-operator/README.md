# How to start

```
if Mongo Db is on localhost:

# For operator
DB_TYPE=local TYPE=operator go run main.go

# For client
DB_TYPE=local TYPE=client go run main.go

if there is a server

PASSWORD_DB=
IP=
LOGIN_DB=

# For operator
DB_TYPE=server PASSWORD_DB=  IP= LOGIN_DB=  TYPE=operator go run main.go

# For client
DB_TYPE=server TYPE=client CONTRACT_ADDRESS=0x8e221ada68e6e666085b0d287c7ae17fe56af175  PVT_KEY=  go run main.go

```

## All ENV variables

PVT_KEY

CONTRACT_ADDRESS

PASSWORD_DB

IP

LOGIN_DB

DB_TYPE

TYPE

# Hello
There are some frameworks here:

go-gin: https://github.com/gin-gonic/gin

go-gin-swagger: https://github.com/swaggo/gin-swagger/

# DB
Right now we are using: MongoDb as temporary example.

FoundationDb for production later ;)

# Swagger


Here is examples how to comment all: https://swaggo.github.io/swaggo.io/declarative_comments_format/

How to add functions for Swagger:

``` js
// @Description Return History of blocks
// @Produce  application/json
// @Success 200 {array} responses.HistoryResponse
// @Router /getHistoryPart [get]

func ResponseHistory(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H {
      "History": "Array",
  })
}

```

1.  go get -u github.com/swaggo/swag/cmd/swag
2.  swag init
3.  Ready

