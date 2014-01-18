package strongo

import (
	"io"
	"testing"
	"github.com/strongo/benchmarks"
)

func BenchmarkStrongo(b *testing.B) {
	wr := new(benchmarks.DevNullWriter)
	t := new(Template_simple)
    for i := 0; i < b.N; i++ {
		t.Render(wr)
    }
}

type Template_simple struct {

}

func (Template_simple) Render(wr io.Writer) (error){
	wr.Write([]byte("Hello, "))
	wr.Write([]byte("stranger"))
	wr.Write([]byte("!"))
	return nil
}
