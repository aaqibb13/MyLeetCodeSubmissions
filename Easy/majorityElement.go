// Given an array nums of size n, return the majority element.

// The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.

 

// Example 1:

// Input: nums = [3,2,3]
// Output: 3
// Example 2:

// Input: nums = [2,2,1,1,1,2,2]
// Output: 2

func majorityElement(nums []int) int {
    hashMap := map[int]int{}
    count := 0
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
	result := 0
	for i, v := range hashMap {
		frequency := int(math.Floor(float64(len(nums) / 2)))
		if v > frequency {
			result = i
		}
	}
	return result
}
