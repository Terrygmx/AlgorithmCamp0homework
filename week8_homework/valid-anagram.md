### 解题思路
因为都是小写字母，用一个26长度的数组来保存各字母出现的次数即可。
在s中出现的字母，出现一次对应位置+1，t中的对应-1。
如果计数数组有任意位置计数不等于0，则两个字符串不是异位词。

### 代码

```golang
func isAnagram(s string, t string) bool {
    m,n:=len(s),len(t)
    if m != n {return false}
    var cnt [26]int
    for i := range s{
        cnt[s[i]-'a']++
        cnt[t[i]-'a']--
    }
    for _,c := range cnt{
        if c != 0 {return false}
    }
    return true
    
}
```