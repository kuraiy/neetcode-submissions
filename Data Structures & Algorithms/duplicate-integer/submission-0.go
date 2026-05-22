func hasDuplicate(nums []int) bool {
    numdubls := make(map[int]struct{})
    for _, v := range nums {
        if _, exists := numdubls[v]; exists {
            return true
        } else {
            numdubls[v] = struct{}{}
        }
    }

    return false
}
