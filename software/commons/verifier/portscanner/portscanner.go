package portscanner

import (
	"fmt"
	"github.com/anvie/port-scanner"
	"time"
)

// Check server
func RunScanner() {
	ps := portscanner.NewPortScanner("localhost", 2*time.Second, 5)
	for {
		if ps.IsOpen(2000) == true {
			fmt.Println("Server started!")
			return
		}
	}
}
