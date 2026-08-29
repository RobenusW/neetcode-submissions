func majorityElement(nums []int) int {
	elem, count := 0, 0
	for _, num := range nums {
		if count == 0 {
			elem = num
		}

		if elem == num {
			count++
		} else {
			count--
			// On the next iteration, the zero count will be taken care of
		}
	}
	return elem
}
