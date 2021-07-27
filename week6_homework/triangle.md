### 解题思路
三角形的特点是，第i行有i个元素。为了求解方便，还是将dp设为n*n的矩阵，而每一行大于i的设为一个较大值，排除他们的影响。
初始值dp[0][0]就是三角形顶点，把他作为递推起始。
按照题意，从第二层开始，每一层的某个节点dp[i][j]，只能从其上一层索引相同的结点或索引-1的节点中继续求值。递推方程就是：
 dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])
排除边缘情况，即j=0的情况即可。
因此，递推方程为
if j==0:  dp[i][j] = triangle[i][j] + dp[i-1][j]
else: dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])

在最后一层，求该层的最小值，即为答案。
### 代码

```golang
func minimumTotal(triangle [][]int) int {
    m,n:= len(triangle),len(triangle) // 为了求解方便，将dp设为n*n的矩阵
    dp := make([][]int,m)
    MAX := 1<<31
    for i := range dp{
        dp[i] = make([]int, n)
        for j:= i+1;j<n;j++ {
            dp[i][j] = MAX // 每一行大于i的设为一个较大值，排除他们的影响
        }
    }
    dp[0][0] = triangle[0][0]
    for i:=1;i<m;i++ {
        for j:=0;j<=i;j++ {
            if j == 0 {
                dp[i][j] = triangle[i][j] + dp[i-1][j]
            }else{
                dp[i][j] = triangle[i][j] + min(dp[i-1][j],dp[i-1][j-1])
            }
        }
    }
    ans := MAX
    for j:=0;j<n;j++ {
        ans = min(ans,dp[m-1][j])
    }
    return ans
}

func min(a,b int) int {
    if a<b {return a}
    return b
}
```