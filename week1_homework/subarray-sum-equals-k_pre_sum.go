func subarraySum(nums []int, k int) int {
    ans, pre := 0, 0
    // m 记录了前缀和情况
    m := map[int]int{}
    // 算上自己
    m[0] = 1
    for i := 0; i < len(nums); i++ {
        pre += nums[i]
        if _, ok := m[pre - k]; ok {
            ans += m[pre - k]
        }
        m[pre] += 1
    }
    return ans
}
