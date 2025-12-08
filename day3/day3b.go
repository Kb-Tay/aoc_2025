package day3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const lim = 12

func Day3b() {
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

		curr_m := findBat(line)
		sum += curr_m
	}

	fmt.Println(sum)
}

func findBat(s []byte) int {

	tmp := make([]byte, 0, lim)
	i := 0

	curr_i := 0

	for ; i < lim; i++ {
		curr_max := byte('0')
		for j := curr_i; j < len(s)-(lim-i-1); j++ {
			if s[j] > curr_max {
				curr_max = s[j]
				curr_i = j + 1
			}
		}

		tmp = append(tmp, curr_max)
	}

	// tmp = append(tmp, remain...)

	v, _ := strconv.Atoi(string(tmp))
	return v
}

// func findBat(s []byte) int {
// 	tmp := make([]byte, 0, lim)
// 	i := 0

// 	for i = 0; i < lim; i++ {
// 		tmp = append(tmp, s[i])
// 	}

// 	for i = lim; i < len(s); i++ {
// 		j := 0
// 		for ; j < len(tmp); j++ {
// 			if s[i] > tmp[j] {
// 				if len(s)-i < lim-j {
// 					continue
// 				}

// 				new_tmp := make([]byte, j+1, lim)
// 				copy(new_tmp, tmp[:j])
// 				new_tmp = append(new_tmp, s[i])
// 				tmp = new_tmp
// 				fmt.Printf("%d %d", int(s[i]-'0'), int(tmp[j]-'0'))
// 				fmt.Println(string(tmp))
// 				break
// 			}
// 		}

// 		if j == len(tmp) && len(tmp) < lim {
// 			fmt.Println(int(s[i] - '0'))
// 			tmp = append(tmp, s[i])
// 		}
// 	}

// 	v, _ := strconv.Atoi(string(tmp))
// 	return v
// }

// func backTrack(choices []byte, s []byte, i int) int {
// 	if len(choices) == 12 {
// 		sum, _ := strconv.Atoi(string(choices))
// 		return sum
// 	}

// 	if i >= len(s) {
// 		return -1
// 	}

// 	new_choice := make([]byte, len(choices), 12)
// 	copy(new_choice, choices)
// 	new_choice = append(new_choice, s[i])
// 	fmt.Println(new_choice)

// 	curr_m := 0

// 	curr_m = max(backTrack(choices, s, i+1), curr_m)
// 	curr_m = max(backTrack(new_choice, s, i+1), curr_m)
// 	// fmt.Println(curr_m)

// 	return curr_m
// }
