/*

  Given a sorted array of distinct integers and a target value, return the index if the target is found. 
  If not, return the index where it would be if it were inserted in order.
  
  Example 1:
 
  Input: nums = [1,3,5,6], target = 5
  Output: 2
 
 */
func searchInsert(nums []int, target int) int {
  index := -1
  for i, v := range nums {
		if target > v {
      index += 1
      continue
		} else if target < v {
			index += 1
			break
		} else if target == v {
			index = i
		}
	}
	if target > nums[len(nums)-1] {
		index += 1
	}
	return index   
}
