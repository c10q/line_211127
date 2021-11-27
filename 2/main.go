package main

import "fmt"

func solution(arr []int) int {
	result := 0
	var aArs [][]int

	up := true
	if arr[0] >= arr[1] {
		up = false
	}

	upCnt := 0
	dwCnt := 0

	for i := 0; i < len(arr)-1; i++ {

		if up {
			if arr[i] < arr[i+1] {
				upCnt += 1
			}

			if arr[i] == arr[i+1] {
				aArs = append(aArs, []int{upCnt, dwCnt})
				upCnt = 0
				dwCnt = 0
			}

			if arr[i] > arr[i+1] {
				up = false
				dwCnt += 1
			}

		} else {
			if arr[i] > arr[i+1] {
				dwCnt += 1
			}

			if arr[i] == arr[i+1] {
				aArs = append(aArs, []int{upCnt, dwCnt})
				upCnt = 0
				dwCnt = 0
			}

			if arr[i] < arr[i+1] {
				up = true
				aArs = append(aArs, []int{upCnt, dwCnt})
				upCnt = 1
				dwCnt = 0
			}
		}
	}

	aArs = append(aArs, []int{upCnt, dwCnt})

	for _, ar := range aArs {
		result += ar[0] * ar[1]
	}

	return result
}

func main() {
	fmt.Println(solution([]int{0, 1, 2, 5, 3, 7}))
	fmt.Println(solution([]int{1, 2, 1, 2, 1}))
	fmt.Println(solution([]int{1, 2, 3, 2, 1}))
	fmt.Println(solution([]int{1, 2, 3, 2, 1, 4, 3, 2, 2, 1}))
}
