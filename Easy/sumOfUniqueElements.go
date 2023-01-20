// You are given an integer array nums. The unique elements of an array are the elements that appear exactly once in the array.

// Return the sum of all the unique elements of nums.

 

// Example 1:

// Input: nums = [1,2,3,2]
// Output: 4
// Explanation: The unique elements are [1,3], and the sum is 4.
// Example 2:

// Input: nums = [1,1,1,1,1]
// Output: 0
// Explanation: There are no unique elements, and the sum is 0.
// Example 3:

// Input: nums = [1,2,3,4,5]
// Output: 15
// Explanation: The unique elements are [1,2,3,4,5], and the sum is 15.


func sumOfUnique(nums []int) int {
    hash := map[int]int{}
    count := 0
      sum := 0
    for i := 0; i < len(nums); i++ {
      if value, ok := hash[nums[i]]; ok {
        value += 1
        hash[nums[i]] = value
      } else {
        count += 1
        hash[nums[i]] = count
        count = 0
      }
    }
    for i, v := range hash {
      if v == 1 {
        sum += i
      }
    }
    return sum
}
