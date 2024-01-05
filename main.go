package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/antoniofrs/experiment-goroutines/src/model"
)

func mapSlice[T, U any](input []T, f func(T) U) []U {
	defer timer("mapSlice")()
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
	defer timer("simpleMapSlice")()
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = f(v)
	}
	return result
}

func timer(funcName string) func() {
    start := time.Now()
    return func() {
        fmt.Printf("%s:\t%v\n", funcName, time.Since(start))
    }
}

func main() {

	students := model.Initialize(1000000)

	go mapSlice(students, model.ToDto)
	go simpleMapSlice(students, model.ToDto)

	time.Sleep(2 * time.Second)
}
