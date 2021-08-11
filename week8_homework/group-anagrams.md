### 解题思路
使用26个字母对应出现频率的26长度的数组作为哈希表的键，然后将同一键结果累加到slice即可。

### 代码

```golang
func groupAnagrams(strs []string) [][]string {
    mapping := make(map[[26]int][]string)
    for _,str := range strs{
        key:= getKey(str)
        mapping[key]=append(mapping[key], str)
    }
    ans := make([][]string,0,len(mapping))
    for _,strs1 := range mapping {
        ans=append(ans, strs1)
    }
    return ans
}
func getKey(str string) (cnt [26]int){
    for _,c := range str{
        cnt[c-'a']++
    }
    return
}
```