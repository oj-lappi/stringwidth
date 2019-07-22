//Very simple package that counts the screen width of a text,
//e.g. for printing cursors (^) under a line pointing at a certain character.
//
//Does not handle combining graphemes or other special characters, only works with monospace runes and tabs.
package stringwidth

import (
	"strings"
)

//Change TabWidth to calculate the output depending on how wide a tab is
var TabWidth = 8

//InSpaces calculates the width in regular space characters between the start of the line and the end of it
func InSpaces(s string) int {
	return strings.Count(s, "") + strings.Count(s, "\t")*(TabWidth-1)
}
