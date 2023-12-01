package bingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	want, got := 6592, Part1().Score
	assert.Equal(t, want, got)
}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = Part1()
	}
}
