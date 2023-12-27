package tidytext

import (
	"slices"
	"strings"
)

type TidyText struct {
	text        string
	indentation int
}

func NewTidyText(text string) *TidyText {
	return &TidyText{text, 4}
}

func (t *TidyText) SetIndentation(indentation int) *TidyText {
	t.indentation = indentation
	return t
}

func (t *TidyText) GetIndentation() int {
	return t.indentation
}

func (t *TidyText) GetText() string {
	return t.text
}

func (t *TidyText) RemoveDoubleNewlines() *TidyText {
	lines := strings.Split(t.text, "\n")
	prevLine := ""
	var newLines []string
	for _, line := range lines {
		if (prevLine == "") && (line == "") {
			continue
		}
		newLines = append(newLines, line)
		prevLine = line
	}
	t.text = strings.Join(newLines, "\n")
	return t
}

func (t *TidyText) TabsToSpaces() *TidyText {
	t.text = strings.Replace(t.text, "\t", strings.Repeat(" ", t.indentation), -1)
	return t
}

func (t *TidyText) TrimEmptyLines() *TidyText {
	lines := strings.Split(t.text, "\n")
	var trimmedLines []string
	for _, line := range lines {
		if strings.Trim(line, " \t") == "" {
			trimmedLines = append(trimmedLines, "")
		} else {
			trimmedLines = append(trimmedLines, line)

		}
	}
	t.text = strings.Join(trimmedLines, "\n")
	return t
}

func (t *TidyText) ReformatIndentation() *TidyText {
	lines := strings.Split(t.text, "\n")
	indents := make([]int, len(lines))
	for i, line := range lines {
		indents[i] = len(line) - len(strings.TrimLeft(line, " \t"))
	}

	uniqueIndents := removeValueFromInts(0, makeUniqueInts(indents))
	slices.Sort(uniqueIndents)

	newIndents := make([]int, len(lines))
	for i := 0; i < len(lines); i++ {
		indent := indents[i]
		if indent == 0 {
			newIndents[i] = 0
			continue
		}
		newIndents[i] = (indexInInts(indent, uniqueIndents) + 1) * t.indentation
	}
	newLines := make([]string, len(lines))
	for i, line := range lines {
		newLines[i] = strings.Repeat(" ", newIndents[i]) + strings.TrimLeft(line, " \t")
	}
	t.text = strings.Join(newLines, "\n")
	return t
}

func (t *TidyText) RemoveLeadingAndTrailingNewlines() *TidyText {
	t.text = strings.Trim(t.text, "\n")
	return t
}

func (t *TidyText) AddTrailingNewLine() *TidyText {
	if (len(t.text) > 0) && (t.text[len(t.text)-1] != '\n') {
		t.text += "\n"
	}
	return t
}

func makeUniqueInts(ints []int) []int {
	var uniqueInts []int
	for _, i := range ints {
		if !intInSlice(i, uniqueInts) {
			uniqueInts = append(uniqueInts, i)
		}
	}
	return uniqueInts
}

func intInSlice(i int, ints []int) bool {
	for _, j := range ints {
		if i == j {
			return true
		}
	}
	return false
}

func indexInInts(i int, ints []int) int {
	for j, k := range ints {
		if i == k {
			return j
		}
	}
	return 0
}

func removeValueFromInts(i int, ints []int) []int {
	var newInts []int
	for _, j := range ints {
		if i != j {
			newInts = append(newInts, j)
		}
	}
	return newInts
}
