package set_matrix_zeroes

func setZeroes(matrix [][]int) {
	var arr [][2]int
	for i, its := range matrix {
		for j, v := range its {
			if v == 0 {
				for k := 0; k < len(its); k++ {
					arr = append(arr, [2]int{i, k})
				}
				for l := 0; l < len(matrix); l++ {
					arr = append(arr, [2]int{l, j})
				}
			}
		}
	}
	for _, ints := range arr {
		matrix[ints[0]][ints[1]] = 0
	}
}

// by key to exercises or problems
func setZeroes2(matrix [][]int) {
	row := make([]bool, len(matrix))
	col := make([]bool, len(matrix[0]))
	for i, r := range matrix {
		for j, v := range r {
			if v == 0 {
				row[i] = true
				col[j] = true
			}
		}
	}
	for i, r := range matrix {
		for j := range r {
			if row[i] || col[j] {
				r[j] = 0
			}
		}
	}
}
