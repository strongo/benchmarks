import jinja2
import logging
from jinja2_keep_module_loader import KeepModuleLoader
from datetime import datetime

logger = logging.getLogger()
logger.setLevel(0)

from data_access import get_authors_with_books


jinja2_settings = {
    'autoescape': False,
    'extensions': [
            #'jinja2.ext.i18n',
            #'tower.template.i18n',
            #'jinja2.ext.autoescape',
            #'jinja2.ext.with_',
        ],
}
JINJA2_TEMPLATES_SOURCE_PATH = 'templates/jinja2/source'

jinja2_environment = jinja2.Environment(
    loader=jinja2.FileSystemLoader(JINJA2_TEMPLATES_SOURCE_PATH),
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
    loader=jinja2.FileSystemLoader(JINJA2_TEMPLATES_SOURCE_PATH),
    #loader=jinja2.ModuleLoader('compiled'),
    #loader=KeepModuleLoader('compiled'),
    **jinja2_settings
    )

from data import users

def test_users_preloaded():
    template = jinja2_environment.get_template('users.html')

    render_counts = 100
    i, s = render_counts, None
    min_spent = None
    logging.info('Warming up Jinja2...')
    template.render(users=users)  # Warm up
    total_started = datetime.now()
    logging.info('Rendering %s times using Jinja2...' % (render_counts,))
    while i > 0:
        single_started = datetime.now()
        s = template.render(users=users)
        single_spent = datetime.now() - single_started
        if not min_spent or single_spent < min_spent:
            min_spent = single_spent
        i -= 1
    total_spent = datetime.now() - total_started
    print s
    logging.info('Minimum render time %s microseconds, Average: %s microseconds' % (min_spent.microseconds, (float(total_spent.microseconds)/render_counts)))
