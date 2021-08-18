### 解题思路
朴素BFS实现思路，从左上角(0,0)出发，对于每一个grid[x][y]等于0的节点，可以考虑其8个方向的子节点。如果任一子节点值为0，且未访问过，可以加入队列，继续BFS。

做完一层BFS，层深+1. 最后到了右下角的节点即grid[n-1][n-1]时，如果该节点取值为0，即可返回层数。

返回-1的情况：
1. 左上角(0,0)为1
2. 不可到达右下角节点
3. 右下角节点为1

### 代码

```python3
class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        n=len(grid)
        if grid[0][0] !=0: return -1
        dx = [0,0,1,1,-1,-1,1,-1]
        dy = [1,-1,1,-1,1,-1,0,0]
        q =[(0,0)]
        visited = set()
        visited.add((0,0))
        level = 1
        def isLegal(x,y):
            return x>=0 and x<n and y>=0 and y < n
        while q:
            # print(q)
            tmp = []
            for ele in q:
                if ele == (n-1,n-1):
                    return level
                i,j = ele
                for k in range(8):
                    x = i+ dx[k]
                    y = j+ dy[k]
                    if isLegal(x,y) and grid[x][y] == 0 and (x,y) not in visited:
                        tmp.append((x,y))
                        visited.add((x,y))
            q = tmp
            level += 1

                

        return -1

```