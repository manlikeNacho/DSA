package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	magazin := make([]int, 26)
	for _, v := range magazine {
		magazin[v-'a']++
	}

	for _, v := range ransomNote {
		magazin[v-'a']--
		if magazin[v-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	canConstruct(
		"a", "b")
}
