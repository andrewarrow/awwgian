package main

import "fmt"
import "time"
import "math/rand"

type InfiniteThing struct {
	space [][][]int8
}

func (t *InfiniteThing) existForInfinity() {
	go t.Expand()
	for {
		fmt.Printf("i am e\n")

		if len(t.space) > 0 {
			i := rand.Intn(len(t.space))
			k := rand.Intn(len(t.space))
			j := rand.Intn(len(t.space))
			fmt.Println(i, j, k, string(t.space[i][j][k]))
		}
		//x := rand.Intn(9223372036854775807) // sometimes you really can't just add 1 more
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
		//go t.Explore(t.space, x)
		time.Sleep(time.Second * 1)
	}
}

func BlankList(size int) []int8 {
	list := []int8{}
	for i := 0; i < size; i++ {
		list = append(list, ri8())
	}
	return list
}
func BlankListOfList(size int) [][]int8 {
	list := [][]int8{}
	for i := 0; i < size; i++ {
		list = append(list, BlankList(size))
	}
	return list
}

func ri8() int8 {
	a := rand.Intn(74) + 48
	if a == 96 {
		return int8(ri8())
	}
	return int8(a)
}
func (t *InfiniteThing) Expand() {
	for {
		if len(t.space) == 0 {
			a := []int8{ri8()}
			b := [][]int8{a}
			t.space = [][][]int8{b}
		} else {
			listOfList := BlankListOfList(len(t.space))
			t.space = append(t.space, append(listOfList, BlankList(len(t.space))))
			for i, ii := range t.space {
				t.space[i] = append(t.space[i], BlankList(len(t.space)))
				for j, _ := range ii {
					t.space[i][j] = append(t.space[i][j], ri8())
				}
			}
		}
		fmt.Println(time.Now(), len(t.space))
		time.Sleep(time.Second * 1)
	}
}

func (t *InfiniteThing) printSpace() {
	for _, ii := range t.space {
		for _, jj := range ii {
			for _, kk := range jj {
				fmt.Printf("%s,", string(kk))
			}
			fmt.Printf("\n")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
}
