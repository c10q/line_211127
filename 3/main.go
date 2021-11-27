package main

import (
	"fmt"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solution(n int, k int) []int {
	var seq [][2]int

	arr := make([][]int, n)
	for i := range arr {
		arr[i] = make([]int, n)
	}

	arr[0][0] = 1
	seq = append(seq, [2]int{0, 0})

	for x := 2; x < (n*n) + 1; x++ {
		distances := make([][]int, n)
		for i := range distances {
			distances[i] = make([]int, n)
		}

		maxD := 0
		var maxDistances [][2]int

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				dMin := n*2
				for _, s := range seq {
					vAbs := Abs(s[0] - i) + Abs(s[1] - j)

					if vAbs < dMin {
						dMin = vAbs
					}
				}
				distances[i][j] = dMin
				if maxD < dMin {
					maxD = dMin
				}
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if distances[j][i] == maxD {
					maxDistances = append(maxDistances, [2]int{j, i})
				}
			}
		}

		i := maxDistances[0][0]
		j := maxDistances[0][1]
		arr[i][j] = x

		seq = append(seq, [2]int{i, j})

		if x == k {
			return []int{i + 1, j + 1}
		}
	}
	return []int{}
}

func main() {
	fmt.Println( solution(5, 12))
	fmt.Println( solution(5, 16))
	fmt.Println( solution(6, 13))
}

