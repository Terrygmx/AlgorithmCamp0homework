### 解题思路
从字符串结尾向前遍历。首先去除结尾的非字母；在遇到字母后，增加ans，直到第二次遇到非字母即可。

### 代码

```cpp
class Solution {
public:
    int lengthOfLastWord(string s) {
        int n = s.size();
        int ans = 0;
        int i = n-1;            
        while( i>=0 && (s[i] > 'z' || s[i] < 'a') && (s[i] > 'Z' || s[i] < 'A')) 
            i--;
        while( i>=0 && ((s[i] <= 'z' && s[i] >= 'a') || (s[i] <= 'Z' && s[i] >= 'A'))) {
            ans++;
            i--;
        }
        
        return ans;
    }
};
```