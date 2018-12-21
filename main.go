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

	// you do believe in the concept of infinity right?
	// most of modern math is based on infinity being
	// a real concept. Something that goes on forever.
	// That's either true of false.
	// Most math people believe true, it's possible for
	// something to go on forever.

	// And yet everyone looking at this loop above knows
	// someone can control-c this program or the whole
	// container/vm/bare metal thing can reboot
	// kill -9 pid there are lots of ways it can stop.
	//
	// So how do you _know_ something can go on forever?
	// Because obviously time started at some point and
	// we are this many second elasped since the begining
	// of time, this is an integer right?
	//
	// And that integer can just keep going up. There's
	// no such thing as max_buffer int overflow in math
	// theory, you just assume well, that never happens,
	// you can always +1 to the int and get a number that
	// is 1 integer higher.
	//
	// But show me an example in the real world where that's
	// possible? Space? Space is infinite and goes on forever?
	// Or it's just a very big integer. Not int32 or int64, more like
	// math.Big huge number but not _actually_ infinite.
	//
	// Before you question whether e could ever die, you have
	// to thing of e like you do all of space and time.

	// long story short: the program can never get here.
}
