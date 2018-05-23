from datetime import datetime

class Author(object):
    __slots__ = 'id', 'first_name', 'last_name', 'top_book_ids', 'top_books'

    def __init__(self, author_id, **kwargs):
        self.id = author_id
        self.first_name = kwargs.get('first_name', None)
        self.last_name = kwargs.get('last_name', None)
        self.top_book_ids = kwargs.get('top_book_ids', None)

    @property
    def full_name(self):
        return '%s %s' % (self.first_name, self.last_name)


class Book(object):
    __slots__ = 'id', 'title', 'author_id',

    def __init__(self, book_id, title, **kwargs):
        self.id = book_id
        self.title = title
        self.author_id = kwargs.get('author_id')

authors = [
    Author(1, first_name='Arthur', last_name='Clark', top_book_ids=[1,2,3]),
]

authors = {
    author.id: author for author in authors
}

books = [
    Book(1, 'Back to starts', author_id=1),
    Book(2, 'Back to starts #2', author_id=1),
    Book(3, 'Back to starts #3', author_id=1),
]

books = {
    book.id: book for book in books
}

class User(object):
    def __init__(self, Id):
        self.Id = Id
        self.Name = "User #%s" % Id
        self.DateOfBirth = datetime.now()

users = [User(i+1) for i in xrange(100)]