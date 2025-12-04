package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day1b() {
	f, err := os.Open("./day1/day1.txt");
	if err != nil {
		fmt.Println(err);
	}
	reader := bufio.NewReader(f);
	lock := 50;
	clicks := 0;
	
	for {
		line, _, err := reader.ReadLine();

		if err != nil {
			break
		}

		dir := string(line[0]);
		num, _ := strconv.Atoi(string(line[1:]));

		if (lock == 0) {
			clicks -= 1
		}

		if (dir == "L") {
			lock -= num;
		} else {
			lock += num;
		}

		// if ans < 0, then add 100 to the ans
		for lock < 0 {
			lock += 100;
			clicks += 1;
		}

		if lock > 99 {
			clicks += lock / 100;
			lock = lock % 100;
		}

		if lock == 0 {
			clicks += 1;
		}
	}	
	fmt.Println(clicks)
}