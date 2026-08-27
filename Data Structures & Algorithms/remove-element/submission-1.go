// Edge Cases
// [], val = x, k = 0 and []
// [0, 1, 2, 20, 1] val = 0, [_, _, _ , _ , _]
	// With zero val all underscore. LEt underscore be 1

// PseudoCode
// First approach
// Iterate throguh the array starting from the last element
// Mantain 2 pointers from the last element
// Iterate one pointer
// When you get to a elem = val
	// Grab the second pointer and bring the second pointer forward to be the value of the first pointer.
	// decrement the second pointer
// Continue until the 0th index is consumed.

// Examples
// Input: nums = [3,2,2,3], val = 3
// [3,2,2,3] i = 2, j = 3
//Output: k = 2, nums = [2,2,_,_]

// Examples
// Input: nums = [0,1,2,2,3,0,4,2], val = 2
// nums = [0,1,2,2,0,3,4,2] i = 4, j = 5
// Output: k = 5, nums = [0,1,3,0,4,_,_,_]

func removeElement(nums []int, val int) int {
	i, j := len(nums) - 1, len(nums) - 1
	notValCount := len(nums)
	for i >= 0 {
		if nums[i] == val {
			nums[i] = nums[j]
			j--
			notValCount--
		}
		i--
	}
	return notValCount
}
