func selectionSort(nums []int) []int {
	// [5, 0]
	for selectedIndex, i := range nums { // 1, 0

		smallestElement := i // 0
		smallestElementIndex := selectedIndex // 1
		for j := selectedIndex + 1; j < len(nums); j++ { // 1, 5
			if smallestElement > nums[j] {  // 5 > 0
				smallestElement = nums[j]
				smallestElementIndex = j 
			}
		}
		nums[selectedIndex] = smallestElement // nums[0] = 0
		nums[smallestElementIndex] = i // nums[1] = 5
	}
	return nums
}


func majorityElement(nums []int) int {
	sorted := selectionSort(nums)
	return sorted[divideByTwoRoundUp(len(sorted)) - 1]
}

func divideByTwoRoundUp(size int) int {
	if size % 2 != 0 {
		return (size / 2) + 1
	}
	return size / 2
}
