### 解题思路
左右双指针，遇到非字母之间跳过，两边都是字母时，交换即可。
因为字符串不能之间修改内容，需要放入字符数组操作。

### 代码

```golang
func reverseOnlyLetters(s string) string {
    n := len(s)
    if n<2 {return s}
    b := make([]byte, n)
    for i:= range s{
        b[i] = byte(s[i])
    }
    l,r := 0, n-1
    for l<r {
        if (b[l] >'Z' || b[l]<'A' ) && (b[l]>'z' || b[l]<'a') {
            l++
            continue
        }
        if (b[r] >'Z' || b[r]<'A' ) && (b[r]>'z' || b[r]<'a') {
            r--
            continue
        }
        b[l],b[r] = b[r],b[l]
        l++
        r--
    }
    return string(b)
}
```