package binarysearch
import "slices"

func SearchInts(list []int, key int) int {
	pos, found := slices.BinarySearch(list,key)
    if found{
        return pos
    }
    return -1
}
