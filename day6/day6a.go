package day6

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day6a() {
	f, err := os.Open("./day6/day6.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	operands := make([][]int, 0)
	var operators []string

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		numStr := strings.Fields(string(line))

		if isOperand(numStr[0]) {
			operators = numStr
			break
		}

		row := make([]int, len(numStr))

		for i, n := range numStr {
			num, _ := strconv.Atoi(n)
			row[i] = num
		}

		operands = append(operands, row)
	}

	sum := make([]int, len(operands[0]))

	for i := 0; i < len(operands); i++ {
		for j := 0; j < len(operands[0]); j++ {
			if i == 0 {
				sum[j] = operands[i][j]
				continue
			}

			sum[j] = operate(sum[j], operands[i][j], operators[j])
		}
	}

	ans := 0
	for _, n := range sum {
		ans += n
	}

	fmt.Println(ans)
}

func isOperand(s string) bool {
	return s == "*" || s == "/" || s == "+" || s == "-"
}

func operate(x int, y int, operand string) int {
	switch operand {
	case "+":
		return x + y
	case "/":
		return x / y
	case "-":
		return x - y
	default:
		return x * y
	}
}
