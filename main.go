package main

import "fmt"
import "time"

var version int = 1

func main() {
	fmt.Println("awwgian ", version)

	for {
		e := Thing{}
		go e.run()
		time.Sleep(time.MINUTE * 1)
		e.firstFork()
		time.Sleep(time.MINUTE * 1)
		e.hmmm()
	}
	// it's just not possible for e to get here
	// this part of the program really doesn't "exist"
	// once it starts up, in fact, maybe it's always been running
	// and has never started up in the first place?
	//
	// it never started, it never ended, it's just always
	// looping, a forever for loop, an infinite loop.
	// infinity never ends so
	//
	// and we won't give it a name
}
