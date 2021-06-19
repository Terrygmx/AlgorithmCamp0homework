func subarraySum(nums []int, k int) int {
    n := len(nums)
    cnt := 0
    for i:= range nums{
        var sum int
        for j:=i;j<n;j++{
            sum+=nums[j]
            if sum == k {
                cnt++
            }
        }
    }
    return cnt
}