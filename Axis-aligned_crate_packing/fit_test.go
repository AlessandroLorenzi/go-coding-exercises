package main

import "testing"

func TestFit1(t *testing.T) {
	c := crate{25, 18, 1}
	b := box{6, 5, 1}
	if fit1(c, b) != 12 {
		t.Errorf("test fit1 failed")
	}

	c = crate{10, 10, 1}
	b = box{1, 1, 1}
	if fit1(c, b) != 100 {
		t.Errorf("test fit1 failed")
	}

	c = crate{5, 5, 1}
	b = box{1, 100, 1}
	if fit1(c, b) != 0 {
		t.Errorf("test fit1 failed")
	}

}

func TestFit2(t *testing.T) {
	c := crate{25, 18, 1}
	b := box{6, 5, 1}
	if fit2(c, b) != 15 {
		t.Errorf("test fit2 failed")
	}
	c = crate{5, 5, 1}
	b = box{6, 5, 1}
	if fit2(c, b) != 0 {
		t.Errorf("test fit2 failed")
	}

}

func TestFitn(t *testing.T) {
	if fitn([]int{3, 4}, []int{1, 2}) != 6 {
		t.Errorf("fitn failed")
	}
	if fitn([]int{123, 456, 789, 1011, 1213, 1415},
		[]int{16, 17, 18, 19, 20, 21}) != 1883443968 {
		t.Errorf("fitn failed")
	}

}
