package main

func search(i int, j int, grid [][]int, mark [][]bool) {

	if i < 0 || i >= len(grid) {
		return
	}

	if j < 0 || j >= len(grid[0]) {
		return
	}

	if mark[i][j] {
		return
	}

	if grid[i][j] == 0 {
		return
	}

	mark[i][j] = true

	search(i-1, j, grid, mark)
	search(i+1, j, grid, mark)
	search(i, j-1, grid, mark)
	search(i, j+1, grid, mark)

	/*
		mark[i][j] = false

		if (i+1 < 4) && (grid[i+1][j] == 1) && (mark[i+1][j] == true) {
			search(i+1, j, grid, mark)
		}

		if ((j+1 < 5) && (grid[i][j+1] == 1) && (mark[i][j+1])) == true {
			search(i, j+1, grid, mark)
		}

		if (i-1 >= 0) && (grid[i-1][j] == 1) && (mark[i-1][j] == true) {
			search(i-1, j, grid, mark)
		}

		if (j-1 >= 0) && (grid[i][j-1] == 1) && (mark[i][j-1] == true) {
			search(i, j-1, grid, mark)
		}
	*/

}

func CountIslands(grid [][]int) int {
	var row int = len(grid)

	if row == 0 {
		return 0
	}
	var col int = len(grid[0])

	var res int = 0

	var mark = [][]bool{
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false}}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if (grid[i][j] == 1) && (mark[i][j] == false) {
				search(i, j, grid, mark)
				res++
			}
		}
	}

	return res
}
