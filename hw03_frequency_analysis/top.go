package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

var re = regexp.MustCompile(`[a-zA-Zа-яА-Я]+(-*[a-zA-Zа-яА-Я]+)*`)

func Top10(s string) []string {
	var result []string
	var wordCountArr []wordCount
	wordsCount := map[string]int{}
	words := re.FindAllString(s, -1)
	if len(words) == 0 {
		return nil
	}
	for i := range words {
		wordsCount[strings.ToLower(words[i])]++
	}

	for key, value := range wordsCount {
		wordCountArr = append(wordCountArr, wordCount{word: key, count: value})
	}

	sort.Slice(wordCountArr, func(i, j int) bool {
		if wordCountArr[i].count == wordCountArr[j].count {
			return wordCountArr[i].word < wordCountArr[j].word
		} else {
			return wordCountArr[i].count > wordCountArr[j].count
		}
	})

	for i, val := range wordCountArr {
		if i >= 10 {
			break
		}
		result = append(result, val.word)
	}

	return result
}
