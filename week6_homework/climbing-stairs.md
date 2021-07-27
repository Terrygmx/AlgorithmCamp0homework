### 解题思路
比如n=3，可以从第一层上两步，也可以从第二层上一步。因此递推方程就是： f(n)=f(n-1)+f(n-2)
把初始的1，2赋值后，从第三层就可以用递推方程求解了。

### 代码

```golang
func climbStairs(n int) int {
    if n <= 2 {return n} // 排除一步，两步情况
    dp := make([]int,n+1)
    dp[1],dp[2] = 1, 2
    for i:=3;i<=n;i++ {
        dp[i]=dp[i-1]+dp[i-2] // 每一层可以从i-i,i-2两个地方踏上来，所以方法是二者之和
    }
    return dp[n]
}
```