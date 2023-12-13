package day

import (
	_ "embed"
	"testing"
)

//go:embed test.txt
var inputTest string

// go test -timeout 30s -run ^TestDayOne$ advent2023/day
func TestDayOne(t *testing.T) {
	input = inputTest
	result := DayOne()
	expected := 155
	if result != expected {
		t.Errorf("DayOne() = %v, want %v", result, expected)
	}
}

// go test -benchmem -run=^$ -bench ^BenchmarkDayOne$ advent2023/day
func BenchmarkDayOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DayOne()
	}
}
