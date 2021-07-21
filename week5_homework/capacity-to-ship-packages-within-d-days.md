### 解题思路
此题也是二分查找的典型题目。关键是如何变成二分的样子。
题意可以按照能否满足条件算，把weights数组变成 [0,0...0,1,1,1...] 的样子。这样就可以使用二分查找出第一个1出现的位置。

### 代码

```golang
var (
    ws []int
)

func shipWithinDays(weights []int, days int) int {
    ws = weights
    sum := 0 
    for _,w := range weights {
        sum += w
    }
    // 按照能否满足条件算，weights数组可以变成 [0,0...0,1,1,1...] 的样子，可以使用二分查找出第一个1出现的位置。
    l,r := 0, sum+1
    for l<r {
        mid := l + (r-l) >> 1 
        if check(mid,days) {
            r = mid
        }else{
            l= mid+1
        }
    }
    return l
}

// 用于检查给定运载能力是否满足条件
func check(w,days int) bool {
    cnt := 0
    for days > 0 && cnt < len(ws) {
        load := w
        for cnt < len(ws) && ws[cnt] <= load {
            load -= ws[cnt]
            cnt++
        }
        days --

    }
    return cnt >= len(ws)
}
```