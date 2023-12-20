package main

import (
	"fmt"
	"hash/fnv"
)

func fnv32(x uint32) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprint(x))) // only string input
	return h.Sum32()
}

func findInput(y, start, end uint32, ch chan uint32, i int) {
	for x := start; x <= end; x++ {
		//fmt.Println("**", i, "** ** x **:", x)
		r := fnv32(x)
		if r == y {
			ch <- x
			close(ch)
		}
	}
}

func main() {
	MaxUint := ^uint32(0)

	ch := make(chan uint32)
	var y uint32   //result
	var thread int //default number

	fmt.Printf("input y:")
	fmt.Scan(&y)
	fmt.Printf("input number of threads:")
	fmt.Scan(&thread)

	// test case
	// y = 1875264835 or 899998 or 1900538226 // x = 3402619870 // thread = 4
	space := MaxUint / uint32(thread)
	start := uint32(0)
	end := space

	for i := 1; i <= thread; i++ {
		//fmt.Println("start and end:", start, end)
		//time.Sleep(5 * time.Second)
		go findInput(y, start, end, ch, i)
		start = end
		end += space
	}

	fmt.Printf("answer : %v\n", <-ch)
}
