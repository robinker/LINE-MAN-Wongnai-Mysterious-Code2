package util

import (
	"strings"
)

func NormalizeString(secret string, meaningLessChar []string) string {
	secret = removeMeaningLessCharacter(secret, meaningLessChar)
	return secret
}

func FindMeaningLessCharacter(secret string) []string {
	dict := collectFrequency(secret)
	maxFrequency := findMaxValue(dict)
	var strToRemove []string
	for key, val := range dict {
		if val == maxFrequency{
			strToRemove = append(strToRemove, key)
		}
	}
	return strToRemove
}

func collectFrequency(secret string) map[string]int {
	dict := make(map[string]int)
	for _, ch := range secret {
		if _, key := dict[string(ch)]; !key {
			dict[string(ch)] = strings.Count(secret , string(ch))
		}
	}
	return dict
}

func findMaxValue(dict map[string]int) int{
	max := -1
	for _, val := range dict {
		if max < val {
			max = val
		}
	}
	return max
}

func removeMeaningLessCharacter(secret string, charToRemove []string) string {
	for _, ch := range charToRemove {
		secret = strings.Replace(secret, ch, "", -1)
	}
	return secret
}
