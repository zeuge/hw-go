package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

const TOP = 10

var normalizeRegex = regexp.MustCompile(`[^a-zA-Zа-яА-ЯёЁ]`)

func readFile(fileName string) (string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func normalizeWord(word string) string {
	word = normalizeRegex.ReplaceAllString(word, "")
	return word
}

func countWords(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		normalized := normalizeWord(word)
		if normalized != "" {
			counts[normalized]++
		}
	}

	return counts
}

func printTopWords(counts map[string]int) {
	type WordCount struct {
		Word  string
		Count int
	}

	wordCounts := make([]WordCount, 0, len(counts))
	for key, value := range counts {
		wordCounts = append(wordCounts, WordCount{Word: key, Count: value})
	}

	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	top := min(TOP, len(wordCounts))
	for i, count := range wordCounts[:top] {
		fmt.Printf("%d. %s - %d\n", i+1, count.Word, count.Count)
	}
}

func main() {
	fileName := flag.String("f", "", "Имя текстового файла для обработки (обязательно)")
	flag.Parse()

	if *fileName == "" {
		fmt.Println("Ошибка: необходимо указать имя файла с помощью флага -f")
		flag.Usage()
		os.Exit(1)
	}

	content, err := readFile(*fileName)
	if err != nil {
		fmt.Printf("Ошибка чтения файла: %v\n", err)
		os.Exit(1)
	}

	counts := countWords(content)
	printTopWords(counts)
}
