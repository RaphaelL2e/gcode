package reshape_the_matrix

func matrixReshape(mat [][]int, r int, c int) [][]int {
	var res [][]int
	k := len(mat)    //行
	j := len(mat[0]) //列

	// 符合二维矩阵规则
	if k*j == r*c {
		m, n := 0, 0
		for i := 0; i < r; i++ {
			var line []int
			for z := 0; z < c; z++ {
				var value int
				if m < k && n < j {
					value = mat[m][n]
					if n == j-1 {
						n = 0
						m++
					} else {
						n++
					}
				}
				line = append(line, value)
			}
			res = append(res, line)
		}
	} else {
		res = mat
	}

	return res
}
