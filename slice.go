package go_utils

import "strings"

func SliceToStr(slice []string) string {
	l := make([]string, 0)

	for _, v := range slice {
		if v != "" {
			l = append(l, v)
		}
	}
	
	return strings.Join(l, ",")
}
