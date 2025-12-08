package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day3a() {
	f, err := os.Open("./day3/day3.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	sum := 0

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		a, _ := strconv.Atoi(string(line[0]))
		b, _ := strconv.Atoi(string(line[1]))

		// assuming the length of string is larger than 2
		if b > a {
			a = b
			b = -1
		}

		for i := 2; i < len(line)-1; i++ {
			tmp, _ := strconv.Atoi(string(line[i]))

			if tmp > a {
				a = tmp
				b = -1
				continue
			}

			if tmp > b {
				b = tmp
			}
		}

		last, _ := strconv.Atoi(string(line[len(line)-1]))

		if last > b {
			b = last
		}

		// fmt.Printf("a: %d b: %d\n", a, b)
		sum += a*10 + b
	}

	fmt.Println(sum)
}
