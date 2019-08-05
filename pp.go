package yz

import "strings"

// TrimSeqTwoBreaks trims seqence of more than two breaks
// and normalize it into two breaks
func TrimSeqTwoBreaks(doc string) string {
	for {
		// pos represents the beginning position of seq
		// consisting of more than 2 breaks
		// Note: pos is the position of byte, not rune
		pos := strings.IndexAny(doc, "\n\n")

		// for the case any seq breaks is not found
		if pos == -1 {
			break
		}

		// cont represents how many breaks continue after p
		cont := 0
		for {
			if []rune(doc)[pos] == "\n" {
			}
		}
	}
}
