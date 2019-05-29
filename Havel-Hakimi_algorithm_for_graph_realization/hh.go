package main

func hh(a answers) bool {
	a = eliminateZeros(a)
	if len(a) == 0 {
		return true
	}
	a = descendingSort(a)
	n := a[0]
	a = a[1:]
	if lenCheck(n, a) {
		return false
	}
	a = frontElimination(n, a)
	return hh(a)
}
