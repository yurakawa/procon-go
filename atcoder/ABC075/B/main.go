package main

import "fmt"


func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var board = make([]string, h)
	dx := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	dy := []int{-1, -1, -1, 0, 0, 1, 1, 1}

	for i := 0; i < h; i++ {
		fmt.Scan(&board[i])
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if board[i][j] == '#'{
				fmt.Print("#")
				continue
			}
			cnt := 0
			for k := 0; k < 8; k++ {
				ni := i + dy[k]
				nj := j + dx[k]
				if 0 <= ni && ni < h && 0 <= nj && nj < w {
					if board[ni][nj] == '#' {
						cnt++
					}
				}

			}
			fmt.Print(cnt)
		}
		fmt.Println()
	}
}
