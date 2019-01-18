package history

import (
	"math/big"
)

var DBInstance = make([]ClientOperation, 0)

type Event struct {
	Sum           string `json:"sum"`
	Who           string `json:"who"`
	Date          int64  `json:"date"`
	OperationType string `json:"operation_type"`
}

type HistoryResponse struct {
	Events []*Event
}

type ClientOperation struct {
	Sum           *big.Int
	Date          int64
	Who           string
	OperationType string
}

func GetAllOperations() []ClientOperation {
	return DBInstance
}

func GetOperationsByDate(date int64) []ClientOperation {
	i := make([]ClientOperation, 0)
	for _, f := range DBInstance {
		if f.Date == date {
			i = append(i, f)
		}
	}
	return i
}

func Log(sum *big.Int, date int64, who string, operationType string) {

	newOperation := ClientOperation{}
	newOperation.Date = date
	newOperation.OperationType = operationType

	if operationType == "Deposit" {
		newOperation.Sum = sum
	} else if operationType == "Transfer" {
		newOperation.Sum = sum
		newOperation.Who = who
	}

	setToDb(newOperation)
}

func setToDb(a ClientOperation) {
	DBInstance = append(DBInstance, a)
}
