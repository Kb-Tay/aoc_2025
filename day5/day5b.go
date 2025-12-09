package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
)

func Day5b() {
	f, err := os.Open("./day5/day5.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	intervals := parse(r)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := merge(intervals)
	count := 0

	for _, m := range merged {
		count += (m[1] - m[0] + 1)
	}

	fmt.Println(count)
}

func merge(intervals [][]int) [][]int {
	i := 0
	new_iter := make([][]int, 0)
	for {
		if i >= len(intervals) {
			break
		}

		start, end := intervals[i][0], intervals[i][1]

		for j := i + 1; j < len(intervals); j++ {
			if intervals[j][0] > end {
				break
			}
			end = max(end, intervals[j][1])
			i++
		}

		new_iter = append(new_iter, []int{start, end})
		i++
	}

	return new_iter
}

func parse(r *bufio.Reader) [][]int {
	intervals := make([][]int, 0)

	for {
		line, _, err := r.ReadLine()

		if err != nil || len(line) == 0 {
			break
		}

		b := bytes.Split(line, []byte("-"))
		start, end := convertInt(b[0]), convertInt(b[1])
		intervals = append(intervals, []int{start, end})
	}

	return intervals
}
