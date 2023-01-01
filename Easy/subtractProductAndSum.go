// Given an integer number n, return the difference between the product of its digits and the sum of its digits.
 

// Example 1:

// Input: n = 234
// Output: 15 
// Explanation: 
// Product of digits = 2 * 3 * 4 = 24 
// Sum of digits = 2 + 3 + 4 = 9 
// Result = 24 - 9 = 15



func subtractProductAndSum(n int) int {
    var product int = 1
	var sum int = 0
	val := strconv.Itoa(n)
	array := strings.Split(val, "")
	for i := 0; i < len(val); i++ {
		convertedVal, _ := strconv.Atoi(array[i])
		product *= convertedVal
		sum += convertedVal
	}
    return product - sum
}
