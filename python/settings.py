# Django settings
import os

SECRET_KEY = 'not so secret'

TEMPLATE_DIRS = (
    os.path.join(os.path.dirname(__file__), 'templates/django'),
)

TEMPLATE_LOADERS = [
    ('django.template.loaders.cached.Loader', (
        'django.template.loaders.filesystem.Loader',
    )),
]