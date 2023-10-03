// package main
//
// import (
// 	"fmt"
// 	"time"
// )
//
// func main() {
//   start := time.Now()
//
//   for i := 0; i < 1000000; i++ {
//
//   }
//
//   end := time.Now()
//
//   fmt.Println(end.Sub(start))
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Print("\x07")
	for {
		// Outputting a bell character
		fmt.Print("\x07")

		// Waiting a random amount of time between 0 to 60 seconds
		waitTime := rand.Intn(61)
        fmt.Printf("Wait time: %d\n", waitTime)
		time.Sleep(time.Duration(waitTime) * time.Second)
	}
}
