package day5

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
)

func Day5a() {
	f, err := os.Open("./day5/day5.txt")
	if err != nil {
		fmt.Println(err)
	}
	r := bufio.NewReader(f)

	intervals, queries := parseInputs(r)

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	slices.Sort(queries)

	merged := mergeIntervals(intervals)
	// for _, test := range new_int {
	// 	fmt.Println(test)
	// }
	m_i := 0
	q_i := 0
	fresh := 0

	for {
		if q_i >= len(queries) || m_i >= len(merged) {
			break
		}

		if queries[q_i] < merged[m_i][0] {
			q_i++
			continue
		}

		if queries[q_i] >= merged[m_i][0] && queries[q_i] <= merged[m_i][1] {
			fresh++
			q_i++
			continue
		}

		if queries[q_i] > merged[m_i][1] {
			m_i++
		}
	}

	fmt.Println(fresh)
}

func mergeIntervals(intervals [][]int) [][]int {
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

func parseInputs(r *bufio.Reader) ([][]int, []int) {
	intervals := make([][]int, 0)
	queries := make([]int, 0)

	for {
		line, _, err := r.ReadLine()

		if err != nil || len(line) == 0 {
			break
		}

		b := bytes.Split(line, []byte("-"))
		start, end := convertInt(b[0]), convertInt(b[1])
		intervals = append(intervals, []int{start, end})
	}

	for {
		line, _, err := r.ReadLine()

		if err != nil {
			break
		}

		queries = append(queries, convertInt(line))
	}

	return intervals, queries
}

func convertInt(s []byte) int {
	v, _ := strconv.Atoi(string(s))
	return v
}
