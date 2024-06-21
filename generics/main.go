package main

import "fmt"

func main() {
	ints := map[string]int64{"first": 12, "second": 23, "third": 34}

	floats := map[string]float64{"first": 12.1, "second": 23.2, "third": 34.3}

	fmt.Printf("Non-Genric Sums: %v and %v\n", SumInts(ints), SumFloats(floats))
}

func SumInts(m map[string]int64) int64 {
	var s int64
	for _, v := range m {
		s += v
	}

	return s
}

func SumFloats(m map[string]float64) float64 {
	var s float64

	for _, v := range m {
		s += v
	}

	return s
}
