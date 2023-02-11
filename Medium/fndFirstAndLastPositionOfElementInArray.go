// 34. Find First and Last Position of Element in Sorted Array

// Given an array of integers nums sorted in non-decreasing order, find the starting and ending position of a given target value.

// If target is not found in the array, return [-1, -1].

// You must write an algorithm with O(log n) runtime complexity.

 

// Example 1:

// Input: nums = [5,7,7,8,8,10], target = 8
// Output: [3,4]
// Example 2:

// Input: nums = [5,7,7,8,8,10], target = 6
// Output: [-1,-1]
// Example 3:

// Input: nums = [], target = 0
// Output: [-1,-1]
 

// Constraints:

// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
// nums is a non-decreasing array.
// -109 <= target <= 109


func searchRange(nums []int, target int) []int {
    hashMap := map[int]int{}
    result := []int{}
    for i := 0; i < len(nums); i++ {
      if nums[i] == target {
        if value, ok := hashMap[i]; ok {
          hashMap[i] = value
        } else {
          hashMap[i] = nums[i]
        }
      }
    }
    for k := range hashMap {
      result = append(result, k)
    }
      if len(result) == 1 {
      result = append(result, result[0])
    }
      if len(result) == 0 {
          return []int{-1, -1}
      } else {
          sort.Ints(result)
      return []int{result[0], result[len(result)-1]}
      }
}
