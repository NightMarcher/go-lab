package basic

import (
	"fmt"
	"time"
)

func DemoTime() {
	now := time.Now()
	fmt.Printf("\n%T => %+v\n", now, now)
}
