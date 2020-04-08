package acronym

import (
	"regexp"
	"strings"
)

// Abbreviate returns the acronym of the given input phrase
func Abbreviate(input string) string {
	var result string
	apostrophes := regexp.MustCompile(`'s`)
	removeds := apostrophes.ReplaceAllString(input, "")
	fields := regexp.MustCompile(`_|\W+| `)
	newStr := fields.ReplaceAllString(removeds, "|")
	pipes := regexp.MustCompile(`\|+`)
	removextrapipes := pipes.ReplaceAllString(newStr, "|")
	//removespaces := regexp.MustCompile(`  `)
	//StrwoSpaces := removespaces.ReplaceAllString(newStr, "")
	//StrwoSpaces := strings.TrimSpace(removextrapipes)
	FinalStr := strings.Split(removextrapipes, "|")
	//splitstring := regexp.MustCompile(` `)
	//finalstr := splitstring.Split(StrwoSpaces, 50)
	for _, value := range FinalStr {
		result = result + strings.ToUpper(string(value[0]))
	}

	return result
}
