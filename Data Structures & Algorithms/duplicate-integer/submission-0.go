func hasDuplicate(nums []int) bool {
	numberSet := make(map[int]struct{}, len(nums))
	
	for _, num := range nums {
		if _, ok := numberSet[num]; !ok {
			numberSet[num] = struct{}{}
		} else {
			return true
		}
	}
	return false
    
}
