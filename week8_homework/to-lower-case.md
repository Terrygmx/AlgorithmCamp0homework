### 解题思路
按字符一个个计算，如果字符是大写字母，则计算一个和A的差作为偏移量。用a加偏移量即可。

### 代码

```cpp
class Solution {
public:
    string toLowerCase(string s) {
        for (int i=0;i<s.size();i++) {
            if(s[i]<='Z' && s[i]>='A')
                s[i] = 'a' + s[i] - 'A';
        }
        return s;
    }
};
```