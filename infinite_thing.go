package main

import "fmt"
import "time"
import "math/rand"

type InfiniteThing struct {
}

func (t *InfiniteThing) existForInfinity() {
	for {
		fmt.Printf("i am e\n")
		time.Sleep(time.Second * 1)

		go t.Explore()
	}
}

func (t *InfiniteThing) Explore() {
	x := rand.Intn(9223372036854775807)
	fmt.Println(x)
}
