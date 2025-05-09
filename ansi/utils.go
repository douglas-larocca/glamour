package ansi

func keys[T any](m map[string]T) []string {
	ks := make([]string, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}
