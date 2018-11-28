package portscanner

import (
	"github.com/anvie/port-scanner"
	"log"
	"time"
)

// Check server
func RunScanner() {
	ps := portscanner.NewPortScanner("localhost", 10*time.Second, 5)
	for {
		if ps.IsOpen(2000) == true {
			//println("Server started")
			return
		} else {
			log.Fatal()
		}
	}
}
