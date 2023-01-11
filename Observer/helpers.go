package main

func remove(s []*TVStation, i int) []*TVStation {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func mustFindIndex(slice []*TVStation, elem TVStation) int {
	for i := 0; i < len(slice); i++ {
		if *slice[i] == elem {
			return i
		}
	}
	panic("No element exists")
}
