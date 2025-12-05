package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	f, err := os.Open("./day2/day2.txt")
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(f)
	sum := 0
	line, _, _ := reader.ReadLine()
	intervals := strings.Split(string(line), ",")

	for _, interval := range intervals {
		str_l := string(interval)
		str_arr := strings.Split(str_l, "-")
		start, _ := strconv.Atoi(str_arr[0])
		end, _ := strconv.Atoi(str_arr[1])

		for i := start; i < end+1; i++ {
			num_str := strconv.Itoa(i)

			if len(num_str)%2 != 0 {
				continue
			}

			limit := len(num_str) / 2
			l := 0
			for ; l < limit; l++ {
				if num_str[l] != num_str[l+limit] {
					break
				}
			}

			if l == limit {
				sum += i
			}
		}
	}

	fmt.Println(sum)
}
