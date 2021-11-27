package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sum(slice []int, start int, end int) int {
	result := 0

	for _, v := range slice[start:end] {
		result += v
	}

	return result
}

func solution(record []string) []int {
	var pSlice_fifo []int
	var pSlice_lifo []int

	saleCnt := 0
	fifo := 0
	lifo := 0

	for _, v := range record {
		slice := strings.Split(v, " ")

		cost, _ := strconv.Atoi(slice[1])
		cnt, _ := strconv.Atoi(slice[2])

		for i := 0; i < cnt; i++ {
			if slice[0] == "P" {
				pSlice_fifo = append(pSlice_fifo, cost)
				pSlice_lifo = append(pSlice_lifo, cost)
			} else {
				saleCnt += 1
				l_fifo := len(pSlice_fifo)
				f := pSlice_fifo[0]
				pSlice_fifo = append(pSlice_fifo[1:l_fifo], []int{}...)
				fifo += f

				l_lifo := len(pSlice_lifo)
				l := pSlice_lifo[l_lifo-1]
				pSlice_lifo = append(pSlice_lifo[0:l_lifo-1], []int{}...)
				lifo += l
			}
		}
	}

	return []int{fifo, lifo}
}

func main() {
	fmt.Println(solution([]string{"P 300 6", "P 500 3", "S 1000 4", "P 600 2", "S 1200 1"}))
}
