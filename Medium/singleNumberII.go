// 137. Single Number II

// Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.

// You must implement a solution with a linear runtime complexity and use only constant extra space.

 

// Example 1:

// Input: nums = [2,2,3,2]
// Output: 3
// Example 2:

// Input: nums = [0,1,0,1,0,1,99]
// Output: 99
 

// Constraints:

// 1 <= nums.length <= 3 * 104
// -231 <= nums[i] <= 231 - 1
// Each element in nums appears exactly three times except for one element which appears once.


func singleNumber(nums []int) int {
  hashMap := map[int]int{}
	count := 0
    var result int
	for i := 0; i < len(nums); i++ {
		if value, ok := hashMap[nums[i]]; ok {
			value += 1
			hashMap[nums[i]] = value
		} else {
			count += 1
			hashMap[nums[i]] = count
			count = 0
		}
	}
	for i, v := range hashMap {
		if v == 1 {
			result = i
		}
	}
  return result   
}
