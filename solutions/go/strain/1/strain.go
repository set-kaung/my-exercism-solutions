package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
func Keep[T any](list []T, filter func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, i := range list {
		if filter(i) {
			result = append(result, i)
		}
	}
	return result
}

func Discard[T any](list []T, filter func(T) bool) []T {
	result := make([]T, 0, len(list))
	for _, i := range list {
		if !filter(i) {
			result = append(result, i)
		}
	}
	return result
}
