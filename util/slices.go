package util

func Intersect[T comparable](a, b []T) []T {
	res := []T{}
	m := make(map[T]bool)

	for _, e := range a {
		m[e] = false
	}

	for _, e := range b {
		ex, ok := m[e]
		if ok && !ex {
			res = append(res, e)
			m[e] = true
		}
	}

	return res
}

func RemoveDups[T comparable](a []T) []T {
	res := []T{}
	m := make(map[T]bool)

	for _, e := range a {
		if !m[e] {
			res = append(res, e)
			m[e] = true
		}
	}

	return res
}
