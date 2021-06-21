package wordcount

import "unicode"

type Frequency map[string]int

func SplitWords(phrase string) []string {
	words := []string{}
	word := ""
	for _, v := range phrase + " " {
		if unicode.IsLetter(v) || unicode.IsSymbol(v) || unicode.IsDigit(v) {
			word += string(unicode.ToLower(v))
		} else if word != "" {
			words = append(words, word)
			word = ""
		}
	}
	return words
}

func WordCount(phrase string) Frequency {
	counter := make(Frequency)
	for _, word := range SplitWords(phrase) {
		counter[word]++
	}
	return counter
}
