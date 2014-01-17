package strongo

import (
	"fmt"
	"testing"
)

func BenchmarkStrongo(b *testing.B) {
    for i := 0; i < b.N; i++ {
        fmt.Sprintf("hello")
    }
}