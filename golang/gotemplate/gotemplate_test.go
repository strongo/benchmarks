package gotemplate

import (
	//"os"
	"github.com/strongo/benchmarks"
	"testing"
	"text/template"
)

func BenchmarkGoTemplateSimple(b *testing.B) {

	t, _ := template.New("simple").Parse("Hello, {{.}}!")

	wr := new(benchmarks.DevNullWriter)

    for i := 0; i < b.N; i++ {
		t.ExecuteTemplate(wr, "simple", "stranger")
    }
}
