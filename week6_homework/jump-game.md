### 贪心解法
尽可能到达最远位置（贪心）。
如果能到达某个位置，那一定能到达它前面的所有位置。

### 代码
```go
func canJump(nums []int) bool {
    var maxLeap int
    max := func (a,b int) int {
        if a > b {return a}
        return b
    }
    for i,v := range nums {
        if i == len(nums)-1 {
            return true
        }
        if v > 0 {
            maxLeap = max(maxLeap,v+i)
        }else{
            if maxLeap <= i {
                return false
            }
        }
        
    }
    return true
}
```