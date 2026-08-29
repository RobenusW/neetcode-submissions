func majorityElement(nums []int) int {
	count := 0
	elem := 0
	for i, num := range nums {
		hasIterated := i != 0 

		if elem == num && hasIterated {
			count++
		} else {
			count--
			if count < 1 {
				elem = num
				count = 1
			}
		}
	}
	return elem
}
