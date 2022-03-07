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
