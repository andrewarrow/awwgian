package main

import "fmt"
import "time"
import "math/rand"

type InfiniteThing struct {
	space [][][]int8
}

func (t *InfiniteThing) existForInfinity() {
	for {
		fmt.Printf("i am e\n")

		x := rand.Intn(9223372036854775807) // sometimes you really can't just add 1 more
		// and yet exploring 9223372036854775807 will
		// take "forever" and therefore is infinity
		// even though 9223372036854775808 is boom, crash.
		// So which is it? Is infinity just a big finite
		// number like 9223372036854775807? Or, because
		// we can'get to 9223372036854775808 it's by
		// definition NOT infinity?
		//
		// This is what should rock moderm math, re-do all your assumptions about
		// when x goes to infinity and just admit, there is a number like 9223372036854775807
		// in the real world that is max int.
		go t.Explore(t.space, x)
		time.Sleep(time.Second * 100)
	}
}

func (t *InfiniteThing) Explore(scope [][][]int8, x int) {
	i := 0
	for {
		if len(t.space) == 0 {
			a := []int8{0}
			b := [][]int8{a}
			t.space = [][][]int8{b}
		} else {
			bar := []int8{}
			for foo := 0; foo < len(t.space); foo++ {
				bar = append(bar, 0)
			}
			b := [][]int8{bar}
			t.space = append(t.space, b)
			for i, ii := range t.space {
				for j, _ := range ii {
					t.space[i][j] = append(t.space[i][j], 0)
				}
				if i == 0 {
					//t.space = append(t.space, b)
				}
			}
		}
		i++
		if i > x {
			break
		}
		t.printSpace()
		time.Sleep(time.Second * 1)
	}
}

func (t *InfiniteThing) printSpace() {
	for i, ii := range t.space {
		fmt.Printf("|%d|", i)
		for _, jj := range ii {
			for _, kk := range jj {
				fmt.Printf("%d,", kk)
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}
