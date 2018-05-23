from os import environ
import logging
import django
from datetime import datetime
from django.template.loader import get_template
from django.template import Context

environ['DJANGO_SETTINGS_MODULE'] = 'settings'

from django.core.wsgi import get_wsgi_application
application = get_wsgi_application()

logger = logging.getLogger()
logger.setLevel(0)

from data import users

def test_users():

    template = get_template('users.html')

    context = Context({'users': users})

    render_counts = 100
    i = render_counts
    print 'Warming up...'
    template.render(context)
    print '1-st render completed, benchamrk started...'
    total_started = datetime.now()
    while i > 0:
        s = template.render(context)
        i -= 1
    total_spent = datetime.now() - total_started
    print s
    average_in_microseconds = float(total_spent.microseconds)/render_counts
    logging.info('\tAverage %s microseconds' % average_in_microseconds)