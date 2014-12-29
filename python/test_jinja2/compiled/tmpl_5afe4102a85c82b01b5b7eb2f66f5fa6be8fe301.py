from __future__ import division
from jinja2.runtime import LoopContext, TemplateReference, Macro, Markup, TemplateRuntimeError, missing, concat, escape, markup_join, unicode_join, to_string, identity, TemplateNotFound
name = 'authors_page_preloaded.html'

def root(context):
    parent_template = None
    if 0: yield None
    parent_template = environment.get_template('_base.html', 'authors_page_preloaded.html')
    for name, parent_block in parent_template.blocks.iteritems():
        context.blocks.setdefault(name, []).append(parent_block)
    for event in parent_template.root_render_func(context):
        yield event

def block_content(context):
    l_authors = context.resolve('authors')
    if 0: yield None
    yield u'\n\t<ul>\n\t\t'
    l_author = missing
    for l_author in l_authors:
        if 0: yield None
        yield u'\n\t\t\t'
        template = environment.get_template('inc_author_li.html', 'authors_page_preloaded.html')
        for event in template.root_render_func(template.new_context(context.parent, True, locals())):
            yield event
        yield u'\n\t\t'
    l_author = missing
    yield u'\n\t</ul>\n'

def block_h1(context):
    if 0: yield None
    yield u'Top authors with top books'

def block_title(context):
    if 0: yield None
    yield u'Top authors with top books'

blocks = {'content': block_content, 'h1': block_h1, 'title': block_title}
debug_info = '1=8&7=14&9=19&10=22&5=29&3=33'