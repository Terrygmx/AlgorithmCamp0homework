### 解题思路
因为只有小写字母，可以用长度26的数组替代哈希表。
可以第一次遍历字符串时填入数组计数，第二次遍历时找出第一个计数为1的字符。
遍历两次原字符串即可得到答案。

### 代码

```golang
func firstUniqChar(s string) int {
    var dic [26]int
    for _,c := range s{
        dic[c-'a'] ++
    }
    for i,c := range s{
        if dic[c-'a'] == 1 {
            return i
        }
    }
    return -1
}
```