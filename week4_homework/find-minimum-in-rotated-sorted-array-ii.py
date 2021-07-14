class Solution:
    def findMin(self, nums: List[int]) -> int:
        # 查询第一个<= nums[right]的数
        # 范围  [0 .. n-1 ] + [n表示不存在]
        if len(nums)==1:
            return nums[0]
        left = 0
        right = len(nums)-1
        # target = nums[-1]

        while left < right:

            mid = (left + right) >> 1
            print("left:{}, right:{},mid:{}\n".format(left,right,mid))

            if nums[mid] < nums[right]:
                right = mid
            elif nums[mid]> nums[right]:
                left = mid + 1
            else:
                right -=1

       
        return nums[right]