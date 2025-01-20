package lab0

import (
	"fmt"
	"io/ioutil"
	//"log"
	"regexp"
	"sort"
	"strings"
)

// Find the top K most common words in a text document.
// 	path: location of the document
//	numWords: number of words to return (i.e. k)
//	charThreshold: character threshold for whether a token qualifies as a word,
//		e.g. charThreshold = 5 means "apple" is a word but "pear" is not.
// Matching is case insensitive, e.g. "Orange" and "orange" is considered the same word.
// A word comprises alphanumeric characters only. All punctuation and other characters
// are removed, e.g. "don't" becomes "dont".
// You should use `checkError` to handle potential errors.
func topWords(path string, numWords int, charThreshold int) []WordCount {
	// TODO: implement me
	// HINT: You may find the `strings.Fields` and `strings.ToLower` functions helpful
	// HINT: To keep only alphanumeric characters, use the regex "[^0-9a-zA-Z]+"

	//0. Open the document
	doc, err := ioutil.ReadFile(path)
	checkError(err)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//1. Read words "strings.Fields" -> readdocument
	//vars: string[] readdocument / wordcount[] returning
	document := strings.Fields(string(doc))
	wordCounts := make([]WordCount, 0)
	r := regexp.MustCompile("[^0-9a-zA-Z]+")

	//For each element of document
	//2. IF STATEMENT: length >= charThreshold
	//		THEN - store WordCount struct
	//			 - Lower case "strings.ToLower" & remove punctuations
	//			 - INSIDE THE FOR LOOP wordcount[]
	//			   3. IF STATEMENT if present
	//					THEN increment count
	//					OTHERWISE add to the last with count 1
	for i := 0; i < len(document); i++ {
		if len(document[i]) >= charThreshold {
			word := r.ReplaceAllString(strings.ToLower(document[i]), "")
			index := contains(wordCounts, word)
			if index == -1 {
				wc := WordCount{word, 1}
				wordCounts = append(wordCounts, wc)
			} else {
				wordCounts[index].Count++
			}
		}
	}

	//4. Call "sortWordCounts" to sort
	//5. Return the first 'numWords' if the resulting WordCount[]
	//has more than 'numWords' elements.
	sortWordCounts(wordCounts)
	if numWords < len(wordCounts) {
		return wordCounts[:numWords]
	} else {
		return wordCounts
	}
}

// A struct that represents how many times a word is observed in a document
type WordCount struct {
	Word  string
	Count int
}

func (wc WordCount) String() string {
	return fmt.Sprintf("%v: %v", wc.Word, wc.Count)
}

// Helper function to sort a list of word counts in place.
// This sorts by the count in decreasing order, breaking ties using the word.
// DO NOT MODIFY THIS FUNCTION!
func sortWordCounts(wordCounts []WordCount) {
	sort.Slice(wordCounts, func(i, j int) bool {
		wc1 := wordCounts[i]
		wc2 := wordCounts[j]
		if wc1.Count == wc2.Count {
			return wc1.Word < wc2.Word
		}
		return wc1.Count > wc2.Count
	})
}

// Helper function to find the index of a word in word counts
// This returns the index of the word in the wordCounts, if present.
// Otherwise, return 0.
func contains(wordCounts []WordCount, s string) int {
	for i := 0; i < len(wordCounts); i++ {
		if wordCounts[i].Word == s {
			return i
		}
	}
	return -1
}
