package utils

import "strings"

func ContainsInArray(str string, arr []string) bool {
	for _, item := range arr {
		if strings.Contains(str, item) {
			return true
		}
	}
	return false
}
