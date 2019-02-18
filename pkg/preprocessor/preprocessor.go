package preprocessor

import (
	"fmt"
)

var COMMENT_CHAR = ";"
var STRING_CHARS = map[string]bool {
	`"`: true,
	`'`: true,
}


func RemoveComments (lines []string) {
	var source string

	// Go through each line in the source
	for lineNum, line := range lines {

		var inString bool = false
		var charNum int = 0
		var strType string

		// Iterate through every character in the line.
		for charNum, _ = range line {
			char := string(line[charNum])

			// This logic sets inString if we're in a string.
			if (inString) && (char == strType) {
				inString = false
			} else if _, exists := STRING_CHARS[char]; exists {
				strType = char
				inString = true
			}

			// Since we've detected a comment declaration outside a string
			// lets cut it.
			if !inString && (char == COMMENT_CHAR) {
				break
			}
		}

		// Cut out the line and add it to the source
		source += line[0:charNum]

		// Throw if there's unterminated string
		if inString {
			errorMsg := fmt.Sprintf("Unterminated string in line %d", lineNum)
			panic(errorMsg)
		}
	}
}