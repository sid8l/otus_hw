package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile(`[A-Za-zА-Яа-я]+-*[A-Za-zА-Яа-я]*`)

type Word struct {
	name     string
	quantity int
}

type Words []Word

func (w Words) getTopWords(quantity int) []string {
	sort.Slice(w, func(i, j int) bool {
		if w[i].quantity == w[j].quantity {
			return w[i].name < w[j].name
		}
		return w[i].quantity > w[j].quantity
	})
	if quantity > len(w) {
		quantity = len(w)
	}
	words := make([]string, quantity)
	for i := 0; i < quantity; i++ {
		words[i] = w[i].name
	}
	return words
}

func Top10(sourceString string) []string {
	wordsQuantity := make(map[string]int)
	topWords := Words{}

	words := re.FindAllString(strings.ToLower(sourceString), -1)
	for i := range words {
		wordsQuantity[words[i]]++
	}
	for key, value := range wordsQuantity {
		topWords = append(topWords, Word{name: key, quantity: value})
	}
	return topWords.getTopWords(10)
}
