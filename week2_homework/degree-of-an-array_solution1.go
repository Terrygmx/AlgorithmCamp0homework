func findShortestSubArray(nums []int) int {
    // 找出数组的度,众数
    fre := make(map[int]int)
    for _,c := range nums{
        fre[c]++
    }
    d, major := 0 , []int{}
    for v,times := range fre{
        if times > d {
            d = times
            major = []int{v}
        }else if times == d {
            major=append(major, v)
        }
    }
   // fmt.Println(d,major)
    //通过众数使用双指针从原数组中找出最小长度
    n := len(nums)
    minLen := n
    for _,c := range major {
        for i:=0;i<n;i++ {
            if nums[i]== c {
                for j:=n-1;j>=i;j-- {
                    if nums[j]== c {
                        minLen = min(minLen, j-i+1)
                        break
                    }
                }
                break
            }
        }

    }
    return minLen

}

func min(a,b int) int {
    if a<b {return a}
    return b
}