package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6b() {
	f, err := os.Open("./day6/day6.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	operands := make([]string, 0)
	var operators []string

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}

		if isOperand(string(line[0])) {
			operators = strings.Fields(string(line))
			break
		}

		operands = append(operands, string(line))
	}

	pos := make([]int, 0)

	for c := 0; c < len(operands[0]); c++ {
		blank := 0

		for r := 0; r < len(operands); r++ {
			if operands[r][c] == ' ' {
				blank += 1
			}
		}

		if blank == len(operands) {
			pos = append(pos, c)
		}
	}

	pos = append(pos, len(operands[0]))

	start, end := 0, 0
	ans := 0

	for i, p := range pos {
		end = p

		test := extract(operands, operators[i], start, end-1)
		ans += test
		start = end + 1
	}

	fmt.Println(ans)
}

func extract(operands []string, op string, s int, e int) int {
	row_len := len(operands)
	sum := 0

	for i := e; i >= s; i-- {
		num := make([]byte, 0)

		for r := 0; r < row_len; r++ {
			ch := operands[r][i]
			if ch == ' ' {
				continue
			}

			num = append(num, operands[r][i])
		}

		n, _ := strconv.Atoi(string(num))

		if sum == 0 {
			sum = n
		} else {
			sum = operate(sum, n, op)
		}
	}

	return sum
}
