package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

func ConcurrentFrequency(texts []string) FreqMap {
	result := FreqMap{}
	ch := make(chan FreqMap)
	wg := &sync.WaitGroup{}
	for _, i := range texts {
		wg.Add(1)
		go func(s string, channel chan FreqMap, wg *sync.WaitGroup) {
			channel <- Frequency(s)
			wg.Done()
		}(i, ch, wg)
	}
	readerWG := sync.WaitGroup{}
	readerWG.Add(1)
	go func() {
		defer readerWG.Done()
		for val := range ch {
			for l, v := range val {
				result[l] += v
			}
		}
	}()
	wg.Wait()
	close(ch)
	readerWG.Wait()
	return result
}
