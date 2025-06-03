package utils

type Array[T any] struct{}

func (a *Array[T]) DelByIdx(arr []T, idx int) []T {
    return append(arr[:idx], arr[idx+1:]...)
}