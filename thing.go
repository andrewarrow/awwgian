package main

import "fmt"
import "time"

type Thing struct {
	waitFactor int
	canFork    bool
}

func (t *Thing) run() {
	for {
		fmt.Printf("i am e\n")
		time.Sleep(time.Second * time.Duration(t.waitFactor))
	}
}
func (t *Thing) firstFork() {
	fmt.Printf("setting limit on forking off...")
	t.canFork = true
}
