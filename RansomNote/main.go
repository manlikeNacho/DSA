package main

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	//makes an array  of 26 characters for the alphabet.
	magazin := make([]int, 26)
	//Maps magazine word to the array using ASIIC value to identify the index on the array
	for _, v := range magazine {
		//Increments the value by 1
		magazin[v-'a']++
	}

	for _, v := range ransomNote {
		//Decrements by 1 and checks for negative value which indicates error
		magazin[v-'a']--
		if magazin[v-'a'] < 0 {
			return false
		}
	}

	//if it manages to survive our test return true
	return true
}

func main() {
	canConstruct(
		"a", "b")
}
