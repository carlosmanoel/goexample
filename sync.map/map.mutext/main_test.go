package main

import (
	"math/rand"
	"testing"
)

var globalResult MyStruct

func nrand(n int) []int {
	i := make([]int, n)
	for ind := range i {
		i[ind] = rand.Int()
	}
	return i
}

func BenchmarkLoadFound(b *testing.B) {
	nums := nrand(b.N)
	rm := NewSyncMap()
	for _, v := range nums {
		var my MyStruct
		my.ID = v
		rm.Store(v, my)
	}

	var currentResult interface{}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		currentResult, _ = rm.Load(nums[i])
	}
	globalResult = currentResult.(MyStruct)
}
