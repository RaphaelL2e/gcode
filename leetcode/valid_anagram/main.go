package valid_anagram

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for _, ch := range t {
		cnt[ch-'a']--
	}
	for _, c := range cnt {
		if c != 0 {
			return false
		}
	}
	return true
}

// 进阶 如果是unicode
// Unicode 一个字符可能对应多个字节
func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	cnt := map[rune]int{}
	for _, ch := range s {
		cnt[ch]++
	}
	for _, ch := range t {
		cnt[ch]--
		if cnt[ch] < 0 {
			return false
		}
	}
	return true
}
