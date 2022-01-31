/* Example 1:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1
*/

func search(nums []int, target int) int {
    indexVal := -1
    for index, value := range nums {
        if value == target {
            indexVal = index
            break
        }
    }
    if indexVal == -1 {
        return -1
    } else {
        return indexVal
    }
}
