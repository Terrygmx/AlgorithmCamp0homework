### 解题思路
由朴素BFS优化.因为终点右下角是已知的,可以采用双端BFS.

### 代码

```python3
class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        n=len(grid)
        if grid[0][0] !=0 or grid[n-1][n-1] != 0: return -1
        dx = [0,0,1,1,-1,-1,1,-1]
        dy = [1,-1,1,-1,1,-1,0,0]
        q =[(0,0)]
        q2 = [(n-1,n-1)]
        visited = set()
        visited2 = set()
        visited.add((0,0))
        visited2.add((n-1,n-1))
        l1,l2 = 1,0
        def isLegal(x,y):
            return x>=0 and x<n and y>=0 and y < n
        while q and q2:
            # print(q)
            tmp = []
            tmp2=[]
            for ele in q:
                if ele in visited2:
                    return l1+l2
                i,j = ele
                for k in range(8):
                    x = i+ dx[k]
                    y = j+ dy[k]
                    if isLegal(x,y) and grid[x][y] == 0 and (x,y) not in visited:
                        tmp.append((x,y))
                        visited.add((x,y))
            q = tmp
            l1 += 1
            for ele in q2:
                if ele in visited:
                    return l1+l2
                i,j = ele
                for k in range(8):
                    x = i+ dx[k]
                    y = j+ dy[k]
                    if isLegal(x,y) and grid[x][y] == 0 and (x,y) not in visited2:
                        tmp2.append((x,y))
                        visited2.add((x,y))
            q2 = tmp2
            l2+=1

                

        return -1
```