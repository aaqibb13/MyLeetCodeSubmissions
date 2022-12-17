// Given an integer array nums and an integer k, 
// return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

// Example 1:
// Input: nums = [1,2,3,1], k = 3
// Output: true

// Example 2:

// Input: nums = [1,0,1,1], k = 1
// Output: true


// Example 3:

// Input: nums = [1,2,3,1,2,3], k = 2
// Output: false

func containsNearbyDuplicate(nums []int, k int) bool {
    dupes := map[int]int{}
    var result bool
    for i, v := range nums {
      val, ok := dupes[v]
      if ok && i-val <= k {
        result = true
      } else {
        dupes[v] = i
      }
    }
    return result
}
