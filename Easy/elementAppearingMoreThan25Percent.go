// Given an integer array sorted in non-decreasing order, there is exactly one integer in the array that occurs more than 25% of the time, return that integer.

 

// Example 1:

// Input: arr = [1,2,2,6,6,6,6,7,10]
// Output: 6
// Example 2:

// Input: arr = [1,1]
// Output: 1
 

// Constraints:

// 1 <= arr.length <= 104
// 0 <= arr[i] <= 105

func findSpecialInteger(arr []int) int {
    requiredPercent := math.Floor(float64(len(arr) / 4))
    count := 0
    result := 0
    hashMap := map[int]int{}
    for i := 0; i < len(arr); i++ {
      if value, ok := hashMap[arr[i]]; ok {
        value += 1
        hashMap[arr[i]] = value
      } else {
        count += 1
        hashMap[arr[i]] = count
        count = 0
      }
    }
    for i, v := range hashMap {
      if v > int(requiredPercent) {
        result = i
        break
      }
    }
    return result
}
