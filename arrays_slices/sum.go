package main

func Sum(array []int) (sum int) {
	for _, val := range array {
		sum += val
	}
	return
}
