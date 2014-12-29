import time
from data import authors, books


class DataProvider(object):
    __slots__ = 'delay_in_seconds',

    def __init__(self, delay_in_seconds):
        self.delay_in_seconds = delay_in_seconds

    def get_authors(self):
        if self.delay_in_seconds:
            time.sleep(self.delay_in_seconds)
        return authors.values()

    def get_books_by_id(self, book_ids):
        if self.delay_in_seconds:
            time.sleep(self.delay_in_seconds)
        return [self.get_book_by_id(book_id) for book_id in book_ids]

    def get_book_by_id(self, book_id):
        if self.delay_in_seconds:
            time.sleep(self.delay_in_seconds)
        return books.get(book_id, None)

