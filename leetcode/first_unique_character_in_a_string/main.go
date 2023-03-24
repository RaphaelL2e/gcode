package first_unique_character_in_a_string

func firstUniqChar(s string) int {
	m := make(map[int32]int)
	for _, b := range s {
		m[b] += 1
	}

	for i, b := range s {
		if m[b] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar2(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range cnt {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}

func firstUniqChar3(s string) int {
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	for i, ch := range s {
		if pos[ch-'a'] == n {
			pos[ch-'a'] = i
		} else {
			pos[ch-'a'] = n + 1
		}
	}
	ans := n
	for _, p := range pos {
		if p < ans {
			ans = p
		}
	}
	if ans < n {
		return ans
	}
	return -1
}
