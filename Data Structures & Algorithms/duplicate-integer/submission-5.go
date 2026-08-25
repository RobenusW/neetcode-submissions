func hasDuplicate(nums []int) bool {
	numberSet := make(map[int]struct{}, len(nums) + 5 )
	
	for _, num := range nums {
		numberSet[num] = struct{}{}
	}
	return !(len(numberSet) == len(nums))
    
}
