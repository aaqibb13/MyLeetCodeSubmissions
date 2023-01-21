// Given a string s, return true if s is a good string, or false otherwise.

// A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).

 

// Example 1:

// Input: s = "abacbc"
// Output: true
// Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.
// Example 2:

// Input: s = "aaabb"
// Output: false
// Explanation: The characters that appear in s are 'a' and 'b'.
// 'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.
 

// Constraints:

// 1 <= s.length <= 1000
// s consists of lowercase English letters.



func areOccurrencesEqual(s string) bool {
  hashMap := map[string]int{}
  count := 0
  for i := 0; i < len(s); i++ {
    if value, ok := hashMap[string(s[i])]; ok {
      value += 1
      hashMap[string(s[i])] = value
    } else {
      count += 1
      hashMap[string(s[i])] = count
      count = 0
    }
  }
  var result bool = false
  previousVal := hashMap[string(s[0])]
  for _, v := range hashMap {
    if previousVal != v {
      result = false
      break
    } else {
      result = true

    }
  }
return result
}
