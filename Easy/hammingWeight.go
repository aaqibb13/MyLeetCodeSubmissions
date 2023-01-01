/*Write a function that takes an unsigned integer and returns the number of '1' bits it has (also known as the Hamming weight).

Note:

Note that in some languages, such as Java, there is no unsigned integer type. In this case, the input will be given as a signed integer type. It should not affect your implementation, as the integer's internal binary representation is the same, whether it is signed or unsigned.
In Java, the compiler represents the signed integers using 2's complement notation. Therefore, in Example 3, the input represents the signed integer. -3.
 

Example 1:

Input: n = 00000000000000000000000000001011
Output: 3
Explanation: The input binary string 00000000000000000000000000001011 has a total of three '1' bits.
*/

func hammingWeight(num uint32) int {
    var convertedNum string = strconv.FormatUint(uint64(num), 2)
	var count int = 0
	var strVal []string = strings.Split(convertedNum, "")
	for i := 0; i < len(strVal); i++ {
		intVal, _ := strconv.Atoi(strVal[i])
		if intVal == 1 {
			count += 1
		}
	}
    return count
}
