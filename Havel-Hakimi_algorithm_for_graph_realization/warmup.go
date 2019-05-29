package main

import "sort"

type answers = []int

func eliminateZeros(a answers) answers {
	na := answers{}
	for _, answer := range a {
		if answer != 0 {
			na = append(na, answer)
		}
	}
	return na
}

func descendingSort(a answers) answers {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	return a
}

func lenCheck(n int, a answers) bool {
	if n <= len(a) {
		return false
	}
	return true
}

func frontElimination(n int, a answers) answers {
	na := answers{}
	for i, answer := range a {
		if i < n {
			na = append(na, answer-1)
		} else {
			na = append(na, answer)
		}
	}
	return na
}

func compareAnswers(a answers, b answers) bool {
	for i, answer := range a {
		if b[i] != answer {
			return false
		}
	}
	return true
}
