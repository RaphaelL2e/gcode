package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	cnt := [26]int{}
	for _, ch := range magazine {
		cnt[ch-'a']++
	}
	for _, ch := range ransomNote {
		cnt[ch-'a']--
	}
	for _, c := range cnt {
		if c < 0 {
			return false
		}
	}
	return true
}
