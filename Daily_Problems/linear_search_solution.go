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
