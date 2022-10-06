package stringlib

import (
	"regexp"
	"strings"
)

// reference : https://stackoverflow.com/questions/28715556/go-equivalent-to-php-preg-match
func PregMatch(str, regexSyntax string) ([]string, error) {
	res := []string{}
	r, err := regexp.Compile(regexSyntax)
	if err != nil {
		return res, err
	}

	for _, match := range r.FindStringSubmatch(str) {
		res = append(res, match)
	}

	return res, err
}

func FindStringIndex(strs []string, str string, cSensitive bool) int {
	str = strings.Trim(str, " ")
	for i, wrd := range strs {
		wrd = strings.Trim(wrd, " ")
		if !cSensitive {
			wrd = strings.ToLower(wrd)
		}

		if wrd == str {
			return i
		}
	}
	return -1
}
