package main

import (
	"fmt"
	"time"
)

func main() {

	closeSign := make(chan struct{})

	for i := 0; i < 2; i++ {
		go func() {
			select {
			case <-closeSign:
				fmt.Println("closed!")
			}
		}()
	}

	t := time.After(2 * time.Second)
	select {
	case <-t:
		close(closeSign)
	}
	time.Sleep(1 * time.Second)
}
