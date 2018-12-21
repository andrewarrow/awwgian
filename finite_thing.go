package main

import "fmt"
import "time"

type FiniteThing struct {
}

func (t *FiniteThing) existForLimetime() {
	for {
		fmt.Printf("i am ?\n")
		time.Sleep(time.Second * 1)
	}
}
