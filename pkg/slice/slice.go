package slice

import (
	"github.com/elliotchance/pie/v2"
)

func UniqueNoNullStringSlice[T string](s []T) []T {
	uniquedSlice := UniqueSlice(s)
	pie.Of(uniquedSlice).Filter(func(t T) bool {
		return t != ""
	})
	return uniquedSlice
}

func UniqueSlice[T comparable](s []T) []T {
	return pie.Unique(s)
}

func UniqueIntSlice(slice ...int) (newSlice []int) {
	found := make(map[int]bool)
	for _, val := range slice {
		if _, ok := found[val]; !ok {
			found[val] = true
			newSlice = append(newSlice, val)
		}
	}
	return
}

func StringSetToSlice(set map[string]bool) []string {
	result := make([]string, 0, len(set))
	for s, ok := range set {
		if ok {
			result = append(result, s)
		}
	}
	return result
}

func StringSliceToSet(slice []string) map[string]bool {
	result := make(map[string]bool, len(slice))
	for _, s := range slice {
		result[s] = true
	}
	return result
}

func StringToInterface(slice []string) []interface{} {
	if len(slice) == 0 {
		return []interface{}{}
	}
	result := make([]interface{}, len(slice))
	for i := range slice {
		result[i] = slice[i]
	}
	return result
}

// Exist check if value exist in slice
func Exist(slice []string, value string) bool {
	return pie.Of(slice).Any(func(s string) bool { return s == value })
}
