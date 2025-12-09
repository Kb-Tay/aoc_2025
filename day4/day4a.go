package day4

import (
	"bufio"
	"fmt"
	"os"
)

func Day4a() {
	f, err := os.Open("./day4/day4.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	mat := make([][]bool, 0)

	// prepare mattrix
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
	ans := 0

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			count := 0

			if !mat[i][j] {
				continue
			}

			for _, dir := range dirs {

				row := dir[0] + i
				col := dir[1] + j

				if row < 0 || row >= len(mat) || col < 0 || col >= len(mat) {
					continue
				}

				if mat[row][col] {
					count++
				}
			}

			if count < 4 {
				// fmt.Printf("r: %d c: %d\n", i, j)
				ans++
			}
		}
	}

	fmt.Println(ans)

	// for i := 0; i < len(mat); i++ {
	// 	for j := 0; j < len(mat[i]); j++ {
	// 		if mat[i][j] {
	// 			fmt.Print("T")
	// 		} else {
	// 			fmt.Print("F")
	// 		}
	// 	}

	// 	fmt.Print("\n")
	// }
}
