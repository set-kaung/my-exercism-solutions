package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    if n <= 0{
        return -1, errors.New("faield")
    }
	count := 0
    for n != 1{
        if n & 1 != 1{
            n = n/2
        }else{
            n = 3*n + 1
        }
        count += 1
    }
    return count,nil
}
