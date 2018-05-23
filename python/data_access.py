import time
from data import authors as _authors, books as _books


class DataProvider(object):
    __slots__ = 'delay_in_seconds',

    def __init__(self, delay_in_seconds):
        self.delay_in_seconds = delay_in_seconds

    def get_authors(self):
        if self.delay_in_seconds:
            time.sleep(self.delay_in_seconds)
        return _authors.values()

    def get_books_by_id(self, book_ids):
        if self.delay_in_seconds:
            time.sleep(self.delay_in_seconds)
        return [self.get_book_by_id(book_id) for book_id in book_ids]

    def get_book_by_id(self, book_id):
        if self.delay_in_seconds:
            time.sleep(self.delay_in_seconds)
        return _books.get(book_id, None)


def get_authors_with_books():
    data_provider = DataProvider(0)
    authors = data_provider.get_authors()
    for author in authors:
        author.top_books = data_provider.get_books_by_id(author.top_book_ids)
    return authors
