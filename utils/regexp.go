package utils

import (
	"regexp"
	"strconv"
)

func FindAllStringGroups(re *regexp.Regexp, str string) (result []string) {
	for _, m := range re.FindAllStringSubmatch(str, -1) {
		result = append(result, m[1])
	}
	return
}

func FindAllIntGroups(re *regexp.Regexp, str string) (result []int) {
	for _, m := range re.FindAllStringSubmatch(str, -1) {
		n, err := strconv.Atoi(m[1])

		if err != nil {
			panic("No numbers")
		}

		result = append(result, n)
	}
	return
}
