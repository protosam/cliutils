package cliutils

import (
	"bytes"
	"strings"
)

// Wrap wraps a single line of text, identifying words by ' '.
// New lines will be separated by '\n'. Very long words, such as URLs will not be wrapped.
// Leading spaces on a new line are stripped. Trailing spaces are not stripped.
//
// Parameters:
//     str - the string to be word wrapped
//     wrapLength - the column (a column can fit only one character) to wrap the words at, less than 1 is treated as 1
//
// Returns:
//     a line with newlines inserted
//
// Original source: https://github.com/Masterminds/goutils/blob/f1923532a168b8203bfe956d8cd3b17ebece5982/wordutils.go#L71-L73
func Wrap(str string, wrapLength int) string {
	return WrapCustom(str, wrapLength, "", false)
}

// WrapCustom wraps a single line of text, identifying words by ' '.
// Leading spaces on a new line are stripped. Trailing spaces are not stripped.
//
// Parameters:
//     str - the string to be word wrapped
//     wrapLength - the column number (a column can fit only one character) to wrap the words at, less than 1 is treated as 1
//     newLineStr - the string to insert for a new line, "" uses '\n'
//     wrapLongWords - true if long words (such as URLs) should be wrapped
//
// Returns:
//     a line with newlines inserted
//
// Original source: https://github.com/Masterminds/goutils/blob/f1923532a168b8203bfe956d8cd3b17ebece5982/wordutils.go#L88-L150
func WrapCustom(str string, wrapLength int, newLineStr string, wrapLongWords bool) string {

	if str == "" {
		return ""
	}
	if newLineStr == "" {
		newLineStr = "\n" // TODO Assumes "\n" is seperator. Explore SystemUtils.LINE_SEPARATOR from Apache Commons
	}
	if wrapLength < 1 {
		wrapLength = 1
	}

	inputLineLength := len(str)
	offset := 0

	var wrappedLine bytes.Buffer

	for inputLineLength-offset > wrapLength {

		if rune(str[offset]) == ' ' {
			offset++
			continue
		}

		end := wrapLength + offset + 1
		spaceToWrapAt := strings.LastIndex(str[offset:end], " ") + offset

		if spaceToWrapAt >= offset {
			// normal word (not longer than wrapLength)
			wrappedLine.WriteString(str[offset:spaceToWrapAt])
			wrappedLine.WriteString(newLineStr)
			offset = spaceToWrapAt + 1

		} else {
			// long word or URL
			if wrapLongWords {
				end := wrapLength + offset
				// long words are wrapped one line at a time
				wrappedLine.WriteString(str[offset:end])
				wrappedLine.WriteString(newLineStr)
				offset += wrapLength
			} else {
				// long words aren't wrapped, just extended beyond limit
				end := wrapLength + offset
				index := strings.IndexRune(str[end:], ' ')
				if index == -1 {
					wrappedLine.WriteString(str[offset:])
					offset = inputLineLength
				} else {
					spaceToWrapAt = index + end
					wrappedLine.WriteString(str[offset:spaceToWrapAt])
					wrappedLine.WriteString(newLineStr)
					offset = spaceToWrapAt + 1
				}
			}
		}
	}

	wrappedLine.WriteString(str[offset:])

	return wrappedLine.String()

}
