package banner

import (
	"fmt"
	"strings"

	"golang.org/x/term"
)

const chars = "abcdefghijklmnopqrstuvwxyz"

var runes = make(map[rune]struct{})

func init() {
	for _, char := range chars {
		runes[char] = struct{}{}
	}
}

type Banner interface {
	Fill(text string) error
	Print()
}

var _ Banner = &banner{}

// banner wraps a grid of ASCII text to be printed to stdout.
type banner struct {
	chars []rune
	lines [7][]rune
}

func New() Banner {
	return &banner{}
}

func NewFromText(text string) (Banner, error) {
	b := &banner{}
	if err := b.Fill(text); err != nil {
		return nil, err
	}

	return b, nil
}

func (b *banner) Fill(text string) error {
	text = strings.ToLower(text)
	if err := validate(text); err != nil {
		return err
	}

	b.chars = []rune(text)

	// We want to string together the characters and splice them
	for _, char := range b.chars {
		blockChar := blockChars[char]
		appendToLines(&b.lines, blockChar)
	}

	return nil
}

func (b *banner) Print() {
	var width int
	if term.IsTerminal(0) {
		width, _, _ = term.GetSize(0)
	}

	if width == 0 {
		for _, line := range b.lines {
			fmt.Println(string(line))
		}
	} else {
		// We need to line-wrap our banner manually over multiple lines...
		var (
			done     bool
			start    int
			end      = min(width, len(b.lines[0]))
			totalLen = len(b.lines[0])
		)

		for !done {
			if end == len(b.lines[0]) {
				done = true
			}

			// If the last column isn't pure spaces then need to pad it...
			end = updateLineEnd(b.lines, end)

			for _, line := range b.lines {
				fmt.Println(string(line[start:end]))
			}
			fmt.Println()

			start += end - start
			end = min(end+width, totalLen)
		}
	}

}

func appendToLines(lines *[7][]rune, blockChar string) {
	// First we need to remove the leading newline.
	trimmed := strings.TrimLeft(blockChar, "\n")

	// Then split on newlines
	split := strings.Split(trimmed, "\n")

	if len(split) != 7 {
		panic("invalid block character height")
	}

	for i, line := range lines {
		lines[i] = []rune(string(line) + " " + split[i])
	}

}

func updateLineEnd(lines [7][]rune, end int) int {
	var done bool
	for {
		done = true
		for _, line := range lines {
			if line[end-1] == '█' {
				// Need to pad...
				end--
				done = false
				break
			}
		}

		if done {
			break
		}
	}
	return end
}

func validate(text string) error {
	if text == "" {
		return ErrNoText
	}

	for _, char := range []rune(strings.ReplaceAll(text, " ", "")) {
		if _, ok := runes[char]; !ok {
			return unsupportedRuneError(char)
		}
	}

	return nil
}
