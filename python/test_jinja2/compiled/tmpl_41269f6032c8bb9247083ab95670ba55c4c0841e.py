from __future__ import division
from jinja2.runtime import LoopContext, TemplateReference, Macro, Markup, TemplateRuntimeError, missing, concat, escape, markup_join, unicode_join, to_string, identity, TemplateNotFound
name = 'inc_book_li.html'

def root(context):
    l_book = context.resolve('book')
    if 0: yield None
    yield u'<li>\n\t<h3><a href="#book=%s">%s</a></h3>\n</li>' % (
        environment.getattr(l_book, 'id'), 
        environment.getattr(l_book, 'title'), 
    )

blocks = {}
debug_info = '2=9'