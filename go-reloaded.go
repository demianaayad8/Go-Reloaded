package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	content, err := readFile(inputFile)
	if err != nil {
		fmt.Println("failed reading input file:", err)
		return
	}

	content = processText(content)

	outputFile := os.Args[2]
	err = writeFile(outputFile, content)
	if err != nil {
		fmt.Println("failed writing output file:", err)
		return
	}
}

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	content := string(data)
	return content, nil
}

func writeFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}

func hexToDecimal(s string) string {
	n, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return s
	}
	return strconv.FormatInt(n, 10)
}

func binToDecimal(s string) string {
	n, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return s
	}
	return strconv.FormatInt(n, 10)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

func processModifiers(words []string) []string {
	result := []string{}

	for i := 0; i < len(words); i++ {
		currentWord := words[i]

		if len(result) == 0 {
			if currentWord == "(up)" || currentWord == "(low)" || currentWord == "(hex)" ||
				currentWord == "(bin)" || currentWord == "(cap)" ||
				currentWord == "(up," || currentWord == "(low," || currentWord == "(cap," {
				continue
			}

			result = append(result, currentWord)
			continue
		}

		lastIndex := len(result) - 1

		switch currentWord {
		case "(up)":
			result[lastIndex] = strings.ToUpper(result[lastIndex])

		case "(low)":
			result[lastIndex] = strings.ToLower(result[lastIndex])

		case "(hex)":
			result[lastIndex] = hexToDecimal(result[lastIndex])

		case "(bin)":
			result[lastIndex] = binToDecimal(result[lastIndex])

		case "(cap)":
			result[lastIndex] = capitalize(result[lastIndex])

		case "(up,":
			if i+1 >= len(words) {
				continue
			}

			countText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(countText)
			if err != nil {
				continue
			}

			if count > len(result) {
				count = len(result)
			}

			for j := 0; j < count; j++ {
				index := lastIndex - j
				result[index] = strings.ToUpper(result[index])
			}

			i++

		case "(low,":
			if i+1 >= len(words) {
				continue
			}

			countText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(countText)
			if err != nil {
				continue
			}

			if count > len(result) {
				count = len(result)
			}

			for j := 0; j < count; j++ {
				index := lastIndex - j
				result[index] = strings.ToLower(result[index])
			}

			i++

		case "(cap,":
			if i+1 >= len(words) {
				continue
			}

			countText := strings.TrimSuffix(words[i+1], ")")
			count, err := strconv.Atoi(countText)
			if err != nil {
				continue
			}

			if count > len(result) {
				count = len(result)
			}

			for j := 0; j < count; j++ {
				index := lastIndex - j
				result[index] = capitalize(result[index])
			}

			i++

		default:
			result = append(result, currentWord)
		}
	}

	return result
}

func isPunctuation(s string) bool { // Helper function - checks if string is only punctuation marks
	if len(s) == 0 {
		return false
	}
	punctuationMarks := ".,!?:;"

	for _, char := range s {
		if !strings.ContainsRune(punctuationMarks, char) {
			return false // Found a non-punctuation character
		}
	}

	return true // All characters were punctuation
}

func fixPunctuation(words []string) []string {
	result := []string{}
	punctuationMarks := ".,!?:;"

	for _, word := range words {
		if word == "" {
			continue
		}

		// Case 1: punctuation alone, like "," or "." or "!"
		if isPunctuation(word) && len(result) > 0 {
			lastIndex := len(result) - 1
			result[lastIndex] = result[lastIndex] + word
			continue
		}

		// Case 2: word starts with punctuation, like ",because" or ",s"
		if len(word) > 0 && strings.ContainsRune(punctuationMarks, rune(word[0])) {
			punctuation := word[0:1]
			restOfWord := word[1:]

			if len(result) > 0 {
				lastIndex := len(result) - 1
				result[lastIndex] = result[lastIndex] + punctuation
			}

			if restOfWord != "" {
				result = append(result, restOfWord)
			}

			continue
		}

		// Case 3: normal word
		result = append(result, word)
	}

	return result
}

func fixQuotes(words []string) []string {
	result := []string{}
	insideQuote := false

	for i := 0; i < len(words); i++ {
		word := words[i]

		if word == "'" {
			if !insideQuote {
				// Opening quote: attach to NEXT word
				insideQuote = true

				// Check if there's a next word
				if i+1 < len(words) {
					words[i+1] = "'" + words[i+1]
				}
			} else {
				// Closing quote: attach to PREVIOUS word
				insideQuote = false

				// Check if there are previous words
				if len(result) > 0 {
					lastIndex := len(result) - 1
					result[lastIndex] = result[lastIndex] + "'"
				}
			}
			// Skip the quote itself (don't add to result)
			continue
		}

		// Regular word - add it
		result = append(result, word)
	}

	return result
}

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		if currentWord == "a" || currentWord == "A" {
			if len(nextWord) == 0 {
				continue
			}

			firstLetter := strings.ToLower(nextWord[0:1])

			if firstLetter == "a" || firstLetter == "e" || firstLetter == "i" ||
				firstLetter == "o" || firstLetter == "u" || firstLetter == "h" {

				if currentWord == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return words
}

func processText(text string) string {
	words := strings.Fields(text)

	words = processModifiers(words)
	words = fixPunctuation(words)
	words = fixQuotes(words)
	words = fixArticles(words)

	return strings.Join(words, " ")
}
