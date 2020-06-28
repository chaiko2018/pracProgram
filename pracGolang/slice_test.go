package main

import "testing"

var intList []int

const defaultSize = 60000

func ready() {
	intList = make([]int, defaultSize)
}

// スライスの先頭へ要素を追加するベンチマーク（Loop使用）
func BenchmarkIntListLoop(b *testing.B) {
	ready()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		intListLoop()
	}
}

func intListLoop() []int {
	tmpIntList := make([]int, len(intList)+1)
	tmpIntList[0] = -1
	for i := 0; i < len(intList); i++ {
		tmpIntList[i+1] = intList[i]
	}
	return tmpIntList
}

// スライスの先頭へ要素を追加するベンチマーク（append使用）
func BenchmarkIntListAppend(b *testing.B) {
	ready()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		intListAppend()
	}
}

func intListAppend() []int {
	return append([]int{-1}, intList...)
}
