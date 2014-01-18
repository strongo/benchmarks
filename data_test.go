package benchmarks_test

import (
	"bitbucket.com/djingo"
	"bitbucket.com/djingo/tests/data"
	"bitbucket.com/djingo/tests/templates/djingo/code"
	"bytes"
	//"fmt"
	"strconv"
	"testing"
)

func Test_author(t *testing.T) {
	a := init_author(2)
	c := create_context(0)
	code.Render_t_author(c, a)
	//fmt.Println(c.String())
}

func Test_authors(t *testing.T) {
	a := init_authors(2, 2)
	c := create_context(0)
	code.Render_t_authors(c, a)
	//fmt.Println(c.String())
}

func create_context(buf_len int) djingo.DjingoContext {
	return djingo.Context{Buffer: bytes.NewBuffer(make([]byte, buf_len))}
}

func init_author(books_count int) (a data.Author) {
	a = data.Author{FirstName: "John", LastName: "Smith", Books: make([]data.Book, books_count)}
	for i := 1; i <= books_count; i++ {
		a.Books[i-1] = data.Book{Id: i, Title: "Book #" + strconv.Itoa(i)}
	}
	return
}

func init_authors(authros_count int, books_per_author int) (a []data.Author) {
	a = make([]data.Author, authros_count)
	for i := 0; i < authros_count; i++ {
		a[i] = init_author(books_per_author)
	}
	return
}

func Benchmark_authors(b *testing.B) {
	b.StopTimer()
	a := init_authors(100, 10)
	c := create_context(0)
	b.StartTimer() //restart timer
	for i := 0; i < b.N; i++ {
		code.Render_t_authors(c, a)
	}
}
