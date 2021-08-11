### 解题思路
简单做法就是开一个哈希表，用宝石的字符填入哈希表，再遍历一次石头字符串判断即可。
另外，因为字符集只有大写、小写字符，总共也就52个字符。所以，可以用一个长度52的数组来替代系统给的哈希表。

具体来说，'a' 的ascci值是97，‘A'的是65，二者相差32。可以把大写，小写字母分为两组，哈希函数可以设计为： 
字符c减去'A'的差值，对32取模，再加上这个差值整除32得到偏移因子。
偏移因子乘以一组字母大小即26得到偏移量。

### 代码

```golang
func numJewelsInStones(jewels string, stones string) int {
    var dic [52]bool
    for _,c := range jewels{
        ind := (c-'A') % 32 + (c-'A') / 32 * 26
        dic[ind] = true
    }
    ans := 0
    for _,c := range stones{
        ind := (c-'A') % 32 + (c-'A') / 32 * 26
        if dic[ind]{
            ans++
        }
    }
    return ans
}
```