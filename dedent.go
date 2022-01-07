package cliutils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	emptyLines        = regexp.MustCompile(`(?m)(^\s*$|^$)`)
	leftMarginMarkers = regexp.MustCompile(`(?m)^([ ]+|[\t]+|)`)
)

// Dedent identifies the most common leading indention and purges it
func Dedent(s string) string {
	// purge all empty lines and find all valid margins to be ranked
	margins := leftMarginMarkers.FindAllString(emptyLines.ReplaceAllString(s, ""), -1)

	// sort margins, they must be shortest to longest
	sort.Strings(margins)

	// rank margins as they are tested
	winningMargin := ""
	winningMarginRank := 0

	// the margin currently being tested
	testingMargin := ""
	testingRank := 0
	testing_m_pos := 0

	// margin to test next
	next_m_pos := 0

	// iterate and rank all margins
	for m_pos := next_m_pos; m_pos < len(margins); m_pos++ {
		// first entry, already winning
		if m_pos == 0 {
			testingMargin = margins[m_pos]
			testingRank = 1
			continue
		}

		// used to determine if winningMarginRankIncremend should be increased
		var winningMarginRankIncremend bool

		// empty margins must be absolute matches
		if testingMargin == "" && margins[m_pos] == testingMargin {
			winningMarginRankIncremend = true
		}

		// other margins are ranked on prefix
		if testingMargin != "" && strings.HasPrefix(margins[m_pos], testingMargin) {
			winningMarginRankIncremend = true
		}

		// update the next margin position to test
		if margins[m_pos] != testingMargin && next_m_pos == testing_m_pos {
			next_m_pos = m_pos
		}

		// rank either must be incremented OR loop needs to goto next_m_pos
		if winningMarginRankIncremend {
			testingRank++

			if winningMarginRank < testingRank {
				winningMargin = testingMargin
				winningMarginRank = testingRank
			}
		} else {
			// setup ne test margin
			testingMargin = margins[next_m_pos]
			testingRank = 0
			testing_m_pos = next_m_pos

			// move m_pos to next position (-1 is neccessary to beat m_pos++)
			m_pos = next_m_pos - 1
		}
	}

	// winner has been declared, use regex to clean up input
	if winningMargin != "" {
		fmt.Printf("winner: %#v\n", winningMargin)
		s = regexp.MustCompile("(?m)^"+winningMargin).ReplaceAllString(s, "")
	}

	return s
}
