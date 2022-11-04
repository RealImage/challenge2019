package util

import "sync"

func FanOut[T any](partners []T, out ...chan T) {
	go func() {
		for _, p := range partners {
			for _, c := range out {
				c <- p
			}
		}

		for _, c := range out {
			close(c)
		}
	}()
}

func Merge[T any](in ...chan T) chan T {
	wg := sync.WaitGroup{}
	out := make(chan T, len(in))
	wg.Add(len(in))

	for _, c := range in {
		go func(ch chan T) {
			defer wg.Done()
			for t := range ch {
				out <- t
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
