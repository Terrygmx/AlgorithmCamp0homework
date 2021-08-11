func reverseWords(s string) string {
	n := len(s)
	r := []byte(s) 
	reverse(r) // 得到总体逆序数组
	var out []byte // 此为输出结果
	for fast,slow := 0,0; fast < n; fast++ { //slow定义为一个单词的开始，fast是单词的结束
		if slow < n && r[slow] == ' ' { // 跳过空格
			slow++
		} else if fast == n-1 || r[fast+1] == ' ' {
            // 输出非空，则增加分隔符空格
			if len(out) != 0 {
				out = append(out, ' ')
			}
			out = append(out, reverse(r[slow:fast+1])...) // 将单词求逆后得到原来的单词
			slow = fast + 1 //把slow定位到当前单词的下一个字符处
		}
	}
	return string(out)
}
func reverse(b []byte) []byte{
	for start, end := 0, len(b)-1; start < end; start, end = start+1, end-1 {
		b[start], b[end] = b[end], b[start]
	}
    return b
}