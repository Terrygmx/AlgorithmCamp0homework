func subarraySum(nums []int, k int) int {
    ans := 0
    n := len(nums)
    var backtrack func(ind,cur int)
    backtrack = func(ind,cur int) {
        if ind >= n{
            return
        }else if nums[ind] == cur {
            ans++
        }
        backtrack(ind+1,cur-nums[ind])
        
    }
    for i:= range nums{
        backtrack(i,k)
    }
    return ans
}