package kata

import "strings"

func solution(str, ending string) bool {
	return strings.HasSuffix(str, ending)
}

func solution2(str, ending string) bool {
	return len(str) >= len(ending) && str[len(str)-len(ending):] == ending
}

func sol3(str, ending string) bool {
	s := len(str)
	e := len(ending)
	if s < e {
		return false
	}
	return str[s-e:] == ending
}
