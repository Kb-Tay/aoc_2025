package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day2b() {
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
			n := len(num_str)

			for k := 1; k < n/2+1; k++ {
				if len(num_str)%k != 0 {
					continue
				}

				pat := num_str[0:k]
				j := k
				for ; j < len(num_str); j += k {
					s := num_str[j : j+k]
					if s != pat {
						break
					}
				}

				if j == len(num_str) {
					sum += i
					// fmt.Printf("num: %d k: %d\n", i, k)
					break
				}
			}
		}
	}

	fmt.Println(sum)
}
