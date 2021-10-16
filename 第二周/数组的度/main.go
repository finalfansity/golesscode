package main

import (
	"fmt"
	"math"
)

func findShortestSubArray(nums []int) int {
	var m = make(map[int][]int)
	var max int
	for i := range nums {
		if n, ok := m[nums[i]]; ok {
			n[0]++
			n[2] = i
		} else {
			m[nums[i]] = []int{1, i, math.MaxInt64}
		}
	}
	ans := math.MaxInt64
	for _, v := range m {
		tmax, tans := v[0], v[2]-v[1]+1
		if max < tmax || (max == tmax && ans > tans) {
			max, ans = tmax, tans
		}
	}
	if max == 1 {
		return 1
	}
	return ans
}

func main() {
	var a = []int{1, 2, 2, 3, 1, 4, 2}
	fmt.Println(findShortestSubArray(a))
}
