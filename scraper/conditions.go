package main

type Conditions []struct {
	arr []string
	fn  func(string, string) bool
}

func (cond Conditions) IsMet(s1 string) bool {
	for _, condition := range cond {
		for _, s2 := range condition.arr {
			if !condition.fn(s1, s2) {
				return false
			}
		}
	}
	return true
}
