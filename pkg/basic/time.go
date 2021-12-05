package basic

import (
	"fmt"
	"time"
)

// time
func DemoTime() {
	now := time.Now()
	fmt.Printf("\n%T => %+v\n", now, now)
}

// timer
func DemoTimer() {
	fmt.Println()
	timer := time.NewTimer(1 * time.Second)

	go func() {
		<-timer.C
		fmt.Println("Timer fired")
	}()

	stop := timer.Stop()
	if stop {
		fmt.Println("Timer stopped")
	}
	// time.Sleep(1024 * time.Millisecond)
}

// ticker
func DemoTicker() {
	fmt.Println()
	ticker := time.NewTicker(250 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1024 * time.Millisecond)
	// ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
