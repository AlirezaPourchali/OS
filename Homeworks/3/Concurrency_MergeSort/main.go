package main

import "fmt"

func Merge(ldata []int, rdata []int) []int {
	result := make([]int, len(ldata)+len(rdata))
	l, r := 0, 0

	for i := 0; i < cap(result); i++ {
		switch {
		case l >= len(ldata):
			result[i] = rdata[r]
			r++
		case r >= len(rdata):
			result[i] = ldata[l]
			l++
		case ldata[l] < rdata[r]:
			result[i] = ldata[l]
			l++
		default:
			result[i] = rdata[r]
			r++
		}
	}

	return result
}

func MergeSort(data []int, r chan []int) {
	if len(data) == 1 {
		r <- data
		return
	}

	leftChan := make(chan []int)
	rightChan := make(chan []int)
	middle := len(data) / 2

	go MergeSort(data[:middle], leftChan)
	go MergeSort(data[middle:], rightChan)

	ldata := <-leftChan
	rdata := <-rightChan

	// close(leftChan)
	// close(rightChan)
	r <- Merge(ldata, rdata)

}

func main() {
	s := []int{22, 8, 3, 31, 4, 2, 16, 6, 11, 25, 8, 10, 18, 14, 7, 15}

	result := make(chan []int)
	go MergeSort(s, result)

	r := <-result
	fmt.Println("\n** the answer: **")
	for _, v := range r {
		fmt.Printf("%v ", v)
	}
	fmt.Printf("\n")
}
