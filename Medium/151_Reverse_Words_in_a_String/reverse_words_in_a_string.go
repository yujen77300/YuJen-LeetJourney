package reversewordsinastring

import "strings"

func reverseWords(s string) string {
	allValue := strings.Split(s, " ")
	valueSlice := []string{}

	for i := len(allValue) - 1; i >= 0; i-- {
		if allValue[i] != "" {
			valueSlice = append(valueSlice, allValue[i])
		}
	}

	return strings.Join(valueSlice, " ")
}
