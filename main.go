package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/antoniofrs/experiment-goroutines/src/model"
)

func mapSlice[T, U any](input []T, f func(T) U) []U {
	result := make([]U, len(input))

	var wg sync.WaitGroup
	wg.Add(len(input))

	for i, v := range input {
		go func(i int, v T) {
			defer wg.Done()
			result[i] = f(v)
		}(i, v)
	}

	wg.Wait()
	
	return result
}

func simpleMapSlice[T, U any](input []T, f func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func timer() func() {
    start := time.Now()
    return func() {
        fmt.Printf("Execution time: %v ", time.Since(start))
    }
}

func main() {

	students := model.Initialize(1000000)

    defer timer()()

	//mapSlice(students, model.ToDto)
	simpleMapSlice(students, model.ToDto)

}
