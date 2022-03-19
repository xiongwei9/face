package main

import "testing"

func unique[T comparable](arr []T) []T {
	m := make(map[T]bool)
	for _, item := range arr {
		m[item] = true
	}
	res := make([]T, 0, len(m))
	for k := range m {
		res = append(res, k)
	}
	return res
}

func TestGeneric(t *testing.T) {
	testArr := []int64{1, 1, 3, 4, 5, 4}
	testRes := unique(testArr)
	t.Log(testRes)

	testArr2 := []string{"hello", "hi", "hello", "what?"}
	testRes2 := unique(testArr2)
	t.Log(testRes2)
}
