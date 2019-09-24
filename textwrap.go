package textwrap

import (
	"fmt"
	"strings"
)

// Used to wrap text to a particular number of characters per line

// WordWrapResult represents the wrap operation result
type WordWrapResult struct {
	TextGroups []string
	WordCount  int
}

// IsValid returns false, if a word count was stored, but no text groups processed, which means
// the max word count was exceeded.
func (w WordWrapResult) IsValid() bool {
	if len(w.TextGroups) == 0 && w.WordCount == 0 {
		return true
	}
	return len(w.TextGroups) > 0 && w.WordCount > 0
}

// WordWrap wraps the specified text to the given width (num of chars), returning an array element for each
// wrapped text group. If the maxWordCount is one or more, then wrapping will exit, if the word count exceeds
// the maxWordCount value.
func WordWrap(fullText string, width, maxWordCount int) (WordWrapResult, error) {
	result := WordWrapResult{}

	// get "words" determined by one or more whitespaces
	words := strings.Fields(fullText)
	result.WordCount = len(words)

	// if max word count exceeded, then don't continue the wrap operation
	if maxWordCount > 0 && result.WordCount > maxWordCount {
		return result, fmt.Errorf("Exceeded word count: %d > %d", result.WordCount, maxWordCount)
	}

	// bail fast, if possible
	switch result.WordCount {
	case 0:
		return result, nil
	case 1:
		result.TextGroups = []string{words[0]}
		return result, nil
	}

	var groups []string
	var sb strings.Builder
	// count the length of each word, once a word causes the group to exceed
	// the width, then place it in a new group. this only fails, if
	// a single word exceeds the width value
	groupLen := 0
	for idx, word := range words {
		groupLen += len(word) + 1 // +1 for whitespace

		if idx == 0 {
			// add first word to builder
			sb.WriteString(word)
		} else if groupLen > width {
			// add another text group
			groups = append(groups, sb.String())

			sb.Reset()

			// add the word that would have went beyond the width to a new group
			sb.WriteString(word)
			groupLen = len(word)
		} else {
			// re-append the word to the final string
			sb.WriteByte(' ')
			sb.WriteString(word)
		}
	}

	// add the last words to the groups
	if sb.Len() > 0 {
		groups = append(groups, sb.String())
	}

	result.TextGroups = groups

	return result, nil
}
