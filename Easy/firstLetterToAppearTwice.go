// 2351. First Letter to Appear Twice

// Given a string s consisting of lowercase English letters, return the first letter to appear twice.

// Note:

// A letter a appears twice before another letter b if the second occurrence of a is before the second occurrence of b.
// s will contain at least one letter that appears twice.
 

// Example 1:

// Input: s = "abccbaacz"
// Output: "c"
// Explanation:
// The letter 'a' appears on the indexes 0, 5 and 6.
// The letter 'b' appears on the indexes 1 and 4.
// The letter 'c' appears on the indexes 2, 3 and 7.
// The letter 'z' appears on the index 8.
// The letter 'c' is the first letter to appear twice, because out of all the letters the index of its second occurrence is the smallest.
// Example 2:

// Input: s = "abcdd"
// Output: "d"
// Explanation:
// The only letter that appears twice is 'd' so we return 'd'.
 

// Constraints:

// 2 <= s.length <= 100
// s consists of lowercase English letters.
// s has at least one repeated letter.



func repeatedCharacter(s string) byte {
    hashMap := map[byte]int{}
    count := 0
    var result byte
    for i := 0; i < len(s); i++ {
      if value, ok := hashMap[byte(s[i])]; ok {
      value += 1
      if value == 2 {
        result = byte(s[i])
        break
      }
      } else {
        count += 1
        hashMap[byte(s[i])] = count
        count = 0
     }
    }
       return result
}
