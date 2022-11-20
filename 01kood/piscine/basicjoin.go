package piscine

func BasicJoin(elems []string) string {
	var ret string
	for _, str := range elems {
		ret += str
	}
	return ret
}
