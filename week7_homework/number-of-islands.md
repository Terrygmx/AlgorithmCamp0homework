### 解题思路
把每个1的点和相邻的点进行find操作，如果两个点的fa不同，则合并两个点的fa。
由于要统计fa的个数，当发生上面的情况时，结果相当于独立区域-1（原来每个值为1的点看作一个区域）

### 代码

```python3
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        def find(index: int) -> int:
            if fa[index] != index:
                fa[index] = find(fa[index]) # 路径压缩
            return fa[index]
        
        def union(index1: int, index2: int):
            fa[find(index1)] = find(index2)
        fa = []        
        m,n = len(grid),len(grid[0])
        for i in range(m):
            for j in range(n):
                fa.append(i*n + j) # 并查集初始化
        # print(fa)
        dx = [1,-1,0,0]
        dy = [0,0,1,-1]
        ans = 0
        for i in range(m):
            for j in range(n):
                # print(i,j)
                if grid[i][j]=='1':
                    ans += 1
                    for d in range(4):
                        x=i+dx[d]
                        y=j+dy[d]
                        if x >= 0 and x<m and grid[x][j]=='1':
                            if find(i*n+j) != find(x*n+j):
                                union(i*n+j, x*n+j)
                                ans-=1
                        if y >= 0 and y < n and grid[i][y]=='1':
                            if find(i*n+j) != find(i*n + y):
                                ans -= 1
                                union(i*n+j, i*n + y)
        # print(fa)
        return ans
                
        
```