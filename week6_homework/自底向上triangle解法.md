### 解题思路
和自顶向下不同，从底开始，到每一层能得到的最小值就是确定的值了。
从n-1层开始，每个元素dp[i][j]能得到的最小和，必然是其下层符合要求的两个和值的较小一个。不难写出递推方程：
 dp[i][j]= triangle[i][j]+min(dp[i+1][j],dp[i+1][j+1])

这样推到第一层，即得到答案dp[0][0]

另外题目要求的空间复杂度O(N)也可以从此得来，简化dp为一维数组，每次滚动使用即可。
滚动时递推方程为：
dp[j] = triangle[i][j]+min(dp[j],dp[j+1]) // 需要自己维护层数i

最后返回dp[0] 即可

### 代码

```python3
class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        n=len(triangle)
        dp = [[0] * n] * n

        for i in range(n):
            dp[n-1][i] = triangle[n-1][i]

        for i in range(n-2,-1,-1):
            for j in range(i+1):
                dp[i][j]= triangle[i][j]+min(dp[i+1][j],dp[i+1][j+1])

        return dp[0][0]
```

### 滚动数组代码
```python3
class Solution:
    def minimumTotal(self, triangle: List[List[int]]) -> int:
        n=len(triangle)
        dp = [0] * n

        for i in range(n):
            dp[i] = triangle[n-1][i]

        for i in range(n-2,-1,-1):
            for j in range(i+1):
                dp[j]= triangle[i][j]+min(dp[j],dp[j+1])

        return dp[0]	
```