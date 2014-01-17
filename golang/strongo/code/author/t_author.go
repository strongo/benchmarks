package code

import (
	"bitbucket.com/djingo"
	"bitbucket.com/djingo/tests/data"
	"strconv"
)

func Render_t_author(c djingo.DjingoContext, author data.Author) {
	c.WriteString("<div><h2><a href=\"/author/")
	c.WriteString(strconv.Itoa(author.Id))
	c.WriteString("\">")
	c.WriteString(author.FirstName)
	c.WriteString(" ")
	c.WriteString(author.LastName)
	c.WriteString("</a></h2><ul>")
	for _, book := range author.Books {
		Render_t_book(c, book)
	}
	c.WriteString("</ul></div>")
}
