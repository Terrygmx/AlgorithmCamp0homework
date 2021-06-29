func findShortestSubArray(nums []int) int {
    // 定义三个map，保存各数出现频率，各数第一次出现，最后一次出现的位置
    fre,first,last := make(map[int]int),make(map[int]int),make(map[int]int)
    
    for i,c := range nums{
        fre[c]++
        if _,ok := first[c]; !ok {
            first[c]=i
        }
        last[c]=i
    }
    // 找出数组的度,众数
    d, major := 0 , []int{}
    for v,times := range fre{
        if times > d {
            d = times
            major = []int{v}
        }else if times == d {
            major=append(major, v)
        }
    }
    // 求众数的最后一个位置和第一个出现位置的差值，再求最小值
    minLen := len(nums)
    for _,c := range major {
       minLen = min(minLen, last[c]-first[c]+1)

    }
    return minLen

}

func min(a,b int) int {
    if a<b {return a}
    return b
}