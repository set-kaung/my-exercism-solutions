package binarysearch
import "fmt"

func SearchInts(list []int, key int) int {
	l,r := 0, len(list)-1
    for l <= r{
        mid := (l+r) / 2
        fmt.Printf("l: %d, r: %d mid: %d\n", l,r,mid)
        if list[mid] == key{
            return mid
        }else if list[mid] > key{
            r = mid - 1
        }else{
            l = mid + 1
        }
    }
    return -1
}
