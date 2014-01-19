package gotemplate

import (
	"bytes"
	"github.com/strongo/benchmarks"
	"testing"
	"text/template"
)

func getSimpleTemplate() (*template.Template, error) {
	return template.New("simple").Parse("Hello, {{.}}!")
}

func Test_Simple_GoTemplate(t *testing.T) {

	t1, _ := getSimpleTemplate()

	wr := new(bytes.Buffer)

	t1.ExecuteTemplate(wr, "simple", "stranger")

	if wr.String() != "Hello, stranger!" {
		t.Errorf("Unexpected output")
	}

}

func Benchmark_Simple_GoTemplate(b *testing.B) {

	t, _ := getSimpleTemplate()

	wr := new(benchmarks.DevNullWriter)

    for i := 0; i < b.N; i++ {
		t.ExecuteTemplate(wr, "simple", "stranger")
    }
}
