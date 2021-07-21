### 解题思路
此题也是二分查找的典型题目。关键是如何变成二分的样子。
首先吃香蕉速度最慢应该是1，吃0无意义。最快就是数组最大值，这样保证能在piles.length时间内吃完。
题意可以按照能否满足条件算，把吃香蕉速度的序列数组变成 [0,0...0,1,1,1...] 的样子。这样就可以使用二分查找出第一个1出现的位置。


### 代码

```python3
class Solution:

    def minEatingSpeed(self, piles: List[int], h: int) -> int:

        n = max(piles)
        l,r = 1,n+1 # 吃香蕉速度最慢应该是1，吃0无意义，而且check方法还会有除以0的问题
     # 按照能否满足条件算，piles数组可以变成 [0,0...0,1,1,1...] 的样子，可以使用二分查找出第一个1出现的位置。
        while l<r:
            mid = l + (r-l)//2
            if self.check(mid,piles,h):
                r=mid
            else:
                l=mid+1

        return l
#用于检查给速度是否满足条件
    def check(self,eatSpeed,piles,h) -> bool:
        needHours = 0
        for p in piles:
            needHours += (p-1)//eatSpeed + 1

        return needHours <= h

```