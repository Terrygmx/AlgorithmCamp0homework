### 解题思路
冗余连接就是去掉以后不影响原来联通关系的边。对于树来说，所有节点都可以看作根节点的子节点（根自己也算上），这正好符合并查集的特点。如果一个连接是冗余连接，证明这个连接的两个节点的找father操作肯定能找到同一个根节点（即两个节点已经在一棵树里面了）。因此用并查集查找两个节点的father并比较，如果相同则是冗余连接。不同则不是，并且按照此连接信息更新树的联通关系（并查集合并）。

### 代码

```python3
class Solution:
    def findRedundantConnection(self, edges: List[List[int]]) -> List[int]:
        def find(index: int) -> int:
            if fa[index] != index:
                fa[index] = find(fa[index]) # 路径压缩
            return fa[index]
        
        def union(index1: int, index2: int):
            fa[find(index1)] = find(index2)
        # 找出n
        n = 1
        for edge in edges:
            n = max(n,edge[1]+1) # 题目要求1~n，这里需要n为edge中最大值+1

        fa = list(range(n))  # 并查集初始化
        ans = []
        for edge in edges:
            if find(edge[0]) == find(edge[1]): # 如果edge中两个节点的father一样，则认为这个边为冗余连接
                ans = edge
            else: # 否则union合并两个节点
                union(edge[0],edge[1])

        return ans

```