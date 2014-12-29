import jinja2
import logging
from keep_module_loader import KeepModuleLoader
from datetime import datetime

logger = logging.getLogger()
logger.setLevel(0)

from data_access import DataProvider


jinja2_settings = {
    'autoescape': False,
    'extensions': [
            #'jinja2.ext.i18n',
            #'tower.template.i18n',
            #'jinja2.ext.autoescape',
            #'jinja2.ext.with_',
        ],
}

jinja2_environment = jinja2.Environment(
    loader=jinja2.FileSystemLoader('templates'),
    **jinja2_settings
    )


def jinja2_logger(msg):
    logging.info(msg)


jinja2_environment.compile_templates(
    target='compiled',
    extensions=None,
    filter_func=None,
    zip=None,
    log_function=jinja2_logger,
    ignore_errors=False,
    py_compile=False)


jinja2_environment = jinja2.Environment(
    loader=jinja2.FileSystemLoader('templates'),
    #loader=jinja2.ModuleLoader('compiled'),
    #loader=KeepModuleLoader('compiled'),
    **jinja2_settings
    )


def test_preloaded():
    template = jinja2_environment.get_template('authors_page_preloaded.html')

    data_provider = DataProvider(0)

    authors = data_provider.get_authors()

    for author in authors:
        author.top_books = data_provider.get_books_by_id(author.top_book_ids)

    render_counts = 10000
    i, s = render_counts, None
    started = datetime.now()
    while i > 0:
        s = template.render(authors=authors)
        i -= 1
    spent = datetime.now() - started
    logging.info('Rendered %s times in %s microseconds (average %s): %s' % (render_counts, spent.microseconds, float(spent.microseconds)/render_counts, s,))


test_preloaded()