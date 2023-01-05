// There is a function signFunc(x) that returns:

// 1 if x is positive.
// -1 if x is negative.
// 0 if x is equal to 0.
// You are given an integer array nums. Let product be the product of all values in the array nums.

// Return signFunc(product).

// Example 1:
  // Input: nums = [-1,-2,-3,-4,3,2,1]
  // Output: 1
  // Explanation: The product of all values in the array is 144, and signFunc(144) = 1




func arraySign(nums []int) int {
    var product int64 = 1
    num := big.NewInt(int64(product))
    bigNum := big.NewInt(1)
    for i := 0; i < len(nums); i++ {
      bigNum = big.NewInt(int64(nums[i]))
      num = new(big.Int).Mul(num, bigNum)
    }
    sign := num.Sign()
      return sign
}
