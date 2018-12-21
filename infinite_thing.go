package main

import "fmt"
import "time"

type InfiniteThing struct {
	waitFactor int
	canFork    bool
}

func (t *InfiniteThing) existForInfinity() {
	for {
		fmt.Printf("i am e\n")
		time.Sleep(time.Second * 1)
	}
}

func (t *InfiniteThing) init() {
	t.waitFactor = 1
}

func (t *InfiniteThing) run() {
	t.init()
	for {
		fmt.Printf("i am e\n")
		time.Sleep(time.Second * time.Duration(t.waitFactor))
	}
}
func (t *InfiniteThing) firstFork() {
	fmt.Printf("setting limit on forking off...")
	t.canFork = true
}
