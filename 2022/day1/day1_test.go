package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {
	want1, want2 := 68923, 200044
	got1, got2 := Solution()
	assert.Equal(t, want1, got1)
	assert.Equal(t, want2, got2)

}

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = Solution()
	}
}
