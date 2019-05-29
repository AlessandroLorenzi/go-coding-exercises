package main

import (
	"testing"
)

func TestEliminateZeros(t *testing.T) {
	a := answers{1, 2, 4, 0, 2, 0, 0, 0}
	a = eliminateZeros(a)
	if len(a) != 4 {
		t.Errorf("Unexpected slice len (%v instead of 4)", len(a))
	}

	a = answers{1}
	a = eliminateZeros(a)
	if len(a) != 1 {
		t.Errorf("Unexpected slice len (%v instead of 1)", len(a))
	}
}

func TestDescendingSort(t *testing.T) {
	a := answers{1, 5, 4, 1}
	a = descendingSort(a)
	if len(a) != 4 {
		t.Errorf("Unexpected slice len (%v instead of 4)", len(a))
	}

	if a[0] != 5 {
		t.Errorf("Unexpected sorting")
	}
	if a[3] != 1 {
		t.Errorf("Unexpected sorting")
	}
}

func TestLenCheck(t *testing.T) {
	if lenCheck(7, answers{6, 5, 5, 3, 2, 2, 2}) != false {
		t.Errorf("Unexpected result")
	}
	if lenCheck(5, answers{5, 5, 5, 5, 5}) != false {
		t.Errorf("Unexpected result")
	}
	if lenCheck(5, answers{5, 5, 5, 5}) != true {
		t.Errorf("Unexpected result")
	}
	if lenCheck(3, answers{1, 1}) != true {
		t.Errorf("Unexpected result")
	}
	if lenCheck(1, answers{}) != true {
		t.Errorf("Unexpected result")
	}
	if lenCheck(0, answers{}) != false {
		t.Errorf("Unexpected result")
	}
}

func TestFrontElimination(t *testing.T) {
	if !compareAnswers(frontElimination(4, answers{5, 4, 3, 2, 1}), answers{4, 3, 2, 1, 1}) {
		t.Errorf("Unexpected result")
	}
	if !compareAnswers(frontElimination(11, answers{14, 13, 13, 13, 12, 10, 8, 8, 7, 7, 6, 6, 4, 4, 2}), answers{13, 12, 12, 12, 11, 9, 7, 7, 6, 6, 5, 6, 4, 4, 2}) {
		t.Errorf("Unexpected result")
	}
	if !compareAnswers(frontElimination(1, answers{10, 10, 10}), answers{9, 10, 10}) {
		t.Errorf("Unexpected result")
	}
	if !compareAnswers(frontElimination(3, answers{10, 10, 10}), answers{9, 9, 9}) {
		t.Errorf("Unexpected result")
	}
	if !compareAnswers(frontElimination(1, answers{1}), answers{0}) {
		t.Errorf("Unexpected result")
	}
}
