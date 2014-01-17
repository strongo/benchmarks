package golang

type Author struct {
	Id        int
	FirstName string
	LastName  string
	Books     []Book
}

type Book struct {
	Id    int
	Title string
}
