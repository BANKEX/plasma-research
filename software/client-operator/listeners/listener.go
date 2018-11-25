package listeners

import (
	"../db"
	"./event"
	"./storage"
	"fmt"
	"log"
	"strconv"
	"time"
)

func Checker() {
	for {
		if storage.StateForEvent == 1 {
			fmt.Println("\n\n\n")

			event.EventCount++

			event.EventMap[event.EventCount] =
				"Who" + "\n" + storage.Who.String() + "\n" +
					"Amount" + "\n" + storage.Amount.String() + "\n" +
					"BlockHash" + "\n" + storage.EventBlockHash + "\n" +
					"BlockNumber" + "\n" + strconv.Itoa(int(storage.EventBlockNumber))

			storage.StateForEvent = 0

			key := strconv.Itoa(event.EventCount)

			value := event.EventMap[event.EventCount]

			err := db.Event("database").Put([]byte(key),[]byte(value))
			if err != nil {
				log.Println(err)
			}

			fmt.Println(event.EventMap[event.EventCount])
		}

		time.Sleep(time.Second * 0)
	}
}