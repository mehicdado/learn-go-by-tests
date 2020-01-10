package arrayslice

//Sum function return sum of elements of given array
func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}
