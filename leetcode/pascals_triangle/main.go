package pascals_triangle

func generate(numRows int) [][]int {
	res := [][]int{{1}}
	if numRows == 1 {
		return res
	}
	if numRows > 1 {
		for i := 1; i < numRows; i++ {
			previous := res[i-1]
			if l := len(previous); l >= 1 {
				var line []int
				line = append(line, 1)
				for i := 0; i < l-1; i++ {
					line = append(line, previous[i]+previous[i+1])
				}
				line = append(line, 1)
				res = append(res, line)
			}
		}
	}
	return res
}
