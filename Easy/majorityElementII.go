// Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

// Example 1:

// Input: nums = [3,2,3]
// Output: [3]
// Example 2:

// Input: nums = [1]
// Output: [1]
// Example 3:

// Input: nums = [1,2]
// Output: [1,2]

func majorityElement(nums []int) []int {
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
    result := []int{}
    for i, v := range hashMap {
      frequency := int(math.Floor(float64(len(nums) / 3)))
      if v > frequency {
        result = append(result, i)
      }
    }
    return result
}
