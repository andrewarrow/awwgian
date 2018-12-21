package main

import "fmt"
import "time"

type InfiniteThing struct {
}

func (t *InfiniteThing) existForInfinity() {
	for {
		fmt.Printf("i am e\n")
		time.Sleep(time.Second * 1)
	}
}
