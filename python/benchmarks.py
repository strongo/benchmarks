from test_django import test_users as django_test_users
from test_jinja2 import test_users_preloaded as jinja2_test_users_preloaded

django_test_users()
jinja2_test_users_preloaded()