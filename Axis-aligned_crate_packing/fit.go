package main

type crate = struct {
	x int
	y int
	z int
}

type box = struct {
	x int
	y int
	z int
}

func fit1(c crate, b box) int {
	x := c.x / b.x
	y := c.y / b.y
	return int(x) * int(y)
}

func fit2(c crate, b box) int {
	regular := fit1(c, b)
	rotated := fit1(c, box{b.y, b.x, 1})
	if rotated > regular {
		return rotated
	}
	return regular
}

func fitn(c []int, b []int) int {
	max := 0
	for _, box := range permutations(b) {
		fits := 1
		for i := 0; i < len(c); i++ {
			fits = fits * int(c[i]/box[i])
		}
		if fits > max {
			max = fits
		}
	}

	return int(max)
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
