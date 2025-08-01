package bookstore

import (
	"math"
)

const bookPrice = 800

var discounts = map[int]float64{
	1: 0.0,
	2: 0.05,
	3: 0.10,
	4: 0.20,
	5: 0.25,
}

func Cost(books []int) int {
	var counts [5]int
	for _, b := range books {
		if b >= 1 && b <= 5 {
			counts[b-1]++
		}
	}
	memo := make(map[[5]int]int)
	return costHelper(counts, memo)
}

func reduceBooks(books [5]int, indices []int) (newBooks [5]int, ok bool) {
	newBooks = books
	for _, i := range indices {
		if newBooks[i] == 0 {
			return newBooks, false
		}
		newBooks[i]--
	}
	return newBooks, true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func costHelper(books [5]int, memo map[[5]int]int) int {
	if val, ok := memo[books]; ok {
		return val
	}

	done := true
	for _, c := range books {
		if c > 0 {
			done = false
			break
		}
	}
	if done {
		return 0
	}

	minCost := math.MaxInt

	available := []int{}
	for i, c := range books {
		if c > 0 {
			available = append(available, i)
		}
	}

	for size := len(available); size >= 1; size-- {
		combs := combinations(available, size)
		for _, comb := range combs {
			newBooks, ok := reduceBooks(books, comb)
			if !ok {
				continue
			}
			discount := discounts[len(comb)]
			groupCost := int(float64(bookPrice*len(comb)) * (1 - discount))
			totalCost := groupCost + costHelper(newBooks, memo)
			minCost = min(minCost, totalCost)
		}
	}

	memo[books] = minCost
	return minCost
}

func combinations(arr []int, k int) [][]int {
	var res [][]int
	comb := make([]int, k)
	var backtrack func(start, depth int)
	backtrack = func(start, depth int) {
		if depth == k {
			tmp := make([]int, k)
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}
		for i := start; i < len(arr); i++ {
			comb[depth] = arr[i]
			backtrack(i+1, depth+1)
		}
	}
	backtrack(0, 0)
	return res
}
