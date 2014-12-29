from __future__ import division
from jinja2.runtime import LoopContext, TemplateReference, Macro, Markup, TemplateRuntimeError, missing, concat, escape, markup_join, unicode_join, to_string, identity, TemplateNotFound
name = 'inc_author_li.html'

def root(context):
    l_author = context.resolve('author')
    if 0: yield None
    yield u'<li>\n\t<h2><a href="#author=%s">%s</a></h2>\n\t<ul>\n\t\t' % (
        environment.getattr(l_author, 'id'), 
        environment.getattr(l_author, 'full_name'), 
    )
    l_book = missing
    for l_book in environment.getattr(l_author, 'top_books'):
        if 0: yield None
        if 0: dummy(l_author)
        template = environment.get_template('inc_book_li.html', 'inc_author_li.html')
        for event in template.root_render_func(template.new_context(context.parent, True, locals())):
            yield event
    l_book = missing
    yield u'\n\t</ul>\n</li>'

blocks = {}
debug_info = '2=9&4=13'