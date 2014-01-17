package code

import (
	"bitbucket.com/djingo"
	"bitbucket.com/djingo/tests/data"
	"strconv"
)

func Render_t_book(c djingo.DjingoContext, book data.Book) {
	c.WriteString("<div><h3><a href=\"/book/")
	c.WriteString(strconv.Itoa(book.Id))
	c.WriteString(">")
	c.WriteString(book.Title)
	c.WriteString("</a></h3>")
}
