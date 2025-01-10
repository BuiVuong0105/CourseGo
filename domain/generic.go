package domain

import "fmt"

type NumberType interface {
	int | int64 | float64
}

type GenericEntity[T any] struct {
	Next  *GenericEntity[T]
	Value T
}

func mapAny[K any, V NumberType](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))
	for index, value := range arr {
		result[index] = f(value)
	}
	return result
}

func RundCmd() {
	arr := []int{1, 2, 3, 4, 5, 6}
	rs := mapAny[int, int](arr, func(v int) int {
		return v * 2
	})
	fmt.Println(rs)
}
