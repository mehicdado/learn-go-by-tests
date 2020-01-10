package arrayslice

//Sum function return sum of elements of given array.
func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}

	return sum
}

//SumAll will sum elements of each given slice and append results to new slice.
func SumAll(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return
}

//SumAllTails will sum elements for each given slice after excluding first element and
//append the results to new slice
func SumAllTails(numbersToSum ...[]int) (sums []int) {

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
