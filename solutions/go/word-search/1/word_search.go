package wordsearch

import "errors"

var diagonals [4][2]int = [4][2]int{[2]int{1, 1}, [2]int{-1, -1}, [2]int{1, -1}, [2]int{-1, 1}}

func checkHorizontal(i, j int, word []byte, graph [][]byte) *[2]int {
	if i+len(word)-1 < len(graph[j]) {
		end := i + len(word) - 1
		for k := 1; k < len(word); k++ {
			if word[k] != graph[j][i+k] {
				end = -1
				break
			}

		}
		if end != -1 {
			return &[2]int{end, j}
		}
	}
	if i-len(word)+1 >= 0 {
		end := i - len(word) + 1
		for k := 1; k < len(word); k++ {
			if word[k] != graph[j][i-k] {
				end = -1
				break
			}
		}
		if end != -1 {
			return &[2]int{end, j}
		}
	}

	return nil
}

func checkVertical(i, j int, word []byte, graph [][]byte) *[2]int {
	if j+len(word)-1 < len(graph) {
		end := j + len(word) - 1
		for k := 1; k < len(word); k++ {
			if word[k] != graph[j+k][i] {
				end = -1
				break
			}
		}
		if end != -1 {
			return &[2]int{i, end}
		}
	}

	if j-len(word)+1 >= 0 {
		end := j - len(word) + 1
		for k := 1; k < len(word); k++ {
			if word[k] != graph[j-k][i] {
				end = -1
				break
			}
		}
		if end != -1 {
			return &[2]int{i, end}
		}
	}
	return nil
}

func checkDiagonals(i, j int, word []byte, graph [][]byte) *[2]int {
	for _, k := range diagonals {
		dx, dy := k[0], k[1]
		currentX, currentY := i, j
		endX := i + (len(word)-1)*dx
		endY := j + (len(word)-1)*dy
		if !(endY >= 0 && endY < len(graph) && endX >= 0 && endX < len(graph[endY])) {
			continue
		}
		for k := 1; k < len(word); k++ {
			currentX += dx
			currentY += dy
			if word[k] != graph[currentY][currentX] {
				currentX = -1
				currentY = -1
				break
			}
		}
		if currentX != -1 && currentY != -1 {
			return &[2]int{currentX, currentY}
		}
	}
	return nil
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	prefixes := map[byte]*[]string{}
	for _, k := range words {
		p := []byte(k)[0]
		if v, ok := prefixes[p]; ok {
			*v = append(*v, k)
		} else {
			prefixes[p] = &[]string{k}
		}
	}
	graph := [][]byte{}
	for _, x := range puzzle {
		graph = append(graph, []byte(x))
	}
	solution := map[string][2][2]int{}
	for j := 0; j < len(graph); j++ {
		for i := 0; i < len(graph[j]); i++ {
			if list, ok := prefixes[graph[j][i]]; ok {
				for _, w := range *list {
					result := checkHorizontal(i, j, []byte(w), graph)
					if result != nil {
						solution[w] = [2][2]int{[2]int{i, j}, *result}
						continue
					}
					result = checkVertical(i, j, []byte(w), graph)
					if result != nil {
						solution[w] = [2][2]int{[2]int{i, j}, *result}
						continue
					}
					result = checkDiagonals(i, j, []byte(w), graph)
					if result != nil {
						solution[w] = [2][2]int{[2]int{i, j}, *result}
						continue
					}
				}
			}
		}
	}
	if len(solution) != len(words) {
		return nil, errors.New("no solution")
	}
	return solution, nil
}
