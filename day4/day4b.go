package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Day4b() {
	f, err := os.Open("./day4/day4.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	mat := make([][]bool, 0)

	// prepare matrix
	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		row := make([]bool, 0, len(line))

		for _, b := range line {
			row = append(row, b == '@')
		}

		mat = append(mat, row)
	}

	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	queue := make([][]int, 0)
	ans := 0

	// put into initial queue
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if !mat[i][j] {
				continue
			}

			isRemove := canRemove(&mat, dirs, i, j)

			if isRemove {
				mat[i][j] = false
				queue = append(queue, []int{i, j})
			}
		}
	}

	// for _, no := range mat {
	// 	fmt.Println(no)
	// }

	for _, no := range queue {
		fmt.Print(no)
	}

	for {
		if len(queue) == 0 {
			break
		}

		node := queue[0]
		x, y := node[0], node[1]
		queue = queue[1:]

		ans++

		for _, dir := range dirs {
			row := dir[0] + x
			col := dir[1] + y

			if row < 0 || row >= len(mat) || col < 0 || col >= len(mat) {
				continue
			}

			if mat[row][col] && canRemove(&mat, dirs, row, col) {
				mat[row][col] = false
				queue = append(queue, []int{row, col})
			}
		}

		fmt.Printf("\nx: %d y: %d\n", x, y)
		for _, no := range queue {
			fmt.Print(no)
		}
	}

	fmt.Println(ans)
}

func canRemove(mat *[][]bool, dirs [][]int, i int, j int) bool {
	n := len(*mat)
	count := 0

	for _, dir := range dirs {
		row := dir[0] + i
		col := dir[1] + j

		if row < 0 || row >= n || col < 0 || col >= n {
			continue
		}

		if (*mat)[row][col] {
			count++
		}
	}

	return count < 4
}
