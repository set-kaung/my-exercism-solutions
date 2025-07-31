package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("")
	}
	distance := char - 'A'
	maxWidth := 2*(int(distance)+1) - 1
	result := make([][]byte, maxWidth)

	for i := 0; i < maxWidth; i++ {
		temp := make([]byte, maxWidth)
		for j := 0; j < maxWidth; j++ {
			temp[j] = ' '
		}
		result[i] = temp
	}
	center := (maxWidth / 2)
	for i := 0; i <= center; i++ {
		result[i][center-i] = 'A' + byte(i)
		result[i][maxWidth-1-center+i] = 'A' + byte(i)
		result[maxWidth-1-i][center-i] = 'A' + byte(i)
		result[maxWidth-1-i][maxWidth-1-center+i] = 'A' + byte(i)
	}

	sb := strings.Builder{}

	sb.Grow(((maxWidth + 1) * maxWidth))
	for i, x := range result {

		sb.WriteString(string(x))
		if i+1 != len(x) {
			sb.WriteRune('\n')
		}
	}

	return sb.String(), nil
}
