package strongo

import (
	"io"
	"bytes"
	"testing"
//	"github.com/strongo/templates"
	"github.com/strongo/benchmarks"
)

const USER = "stranger"

func Test_Simple_Strongo(t *testing.T) {
	template := new(Template_simple)
	writer := new(bytes.Buffer)
	template.Render(writer, USER)
	s := writer.String()
	if s != "Hello, stranger!" {
		t.Errorf("Unexpected output")
	}
}

func Benchmark_Simple_Strongo(b *testing.B) {
	template := new(Template_simple)
	writer := new(benchmarks.DevNullWriter)
	for i := 0; i < b.N; i++ {
		template.Render(writer, USER)
	}
}

type Template_simple struct {
}

//func (t *Template_simple) Render(writer io.Writer, context interface{}) (error){
//	return t.Render2(writer, string(context))
//}

func (t *Template_simple) Render(writer io.Writer, user string) (error){
	writer.Write([]byte("Hello, "))
	writer.Write([]byte(user))
	writer.Write([]byte("!"))
	return nil
}
