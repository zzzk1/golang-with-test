package array

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}
	return
}

func SumAll(numbersToSum ...[]int) (result []int) {
	for _, nums := range numbersToSum {
		result = append(result, Sum(nums))
	}
	return
}

func SumAllTails(numbersToSum ...[]int) (result []int) {

	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			result = append(result, 0)
		} else {
			tail := nums[len(nums)-1:]
			result = append(result, Sum(tail))
		}
	}
	return
}
