// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
// Return the running sum of nums.

// Example 1:

// Input: nums = [1,2,3,4]
// Output: [1,3,6,10]
// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].
// Example 2:

// Input: nums = [1,1,1,1,1]
// Output: [1,2,3,4,5]
// Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].


func runningSum(nums []int) []int {
    for i := 0; i < len(nums)-1; i++ {
	if i < len(nums) {
		nums[i+1] = nums[i] + nums[i+1]
	} else {
		break
	}
    }
    return nums
}
