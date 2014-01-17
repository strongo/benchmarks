package code

import (
	"bitbucket.com/djingo"
	"bitbucket.com/djingo/tests/data"
)

type t_authors struct {
	// base _base
}

func Render_t_authors(c djingo.DjingoContext, authors []data.Author) {
	c.WriteString("<ul>")
	for _, author := range authors {
		c.WriteString("<li>")
		Render_t_author(c, author)
		c.WriteString("</li>")
	}
	c.WriteString("</ul>")
}
