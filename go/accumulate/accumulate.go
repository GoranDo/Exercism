package accumulate

func Accumulate(inputs []string, f func(string) string) []string {
	cache := []string{}
	for _, input := range inputs {
		cache = append(cache, f(input))
	}
	return cache
}
