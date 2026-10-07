func hasDuplicate(nums []int) bool {
    var m map[int]int = map[int]int{}

    for _, num := range nums {
        _, ok := m[num]
        if ok {
            return true
        }
        m[num]++
    }
    return false
}
