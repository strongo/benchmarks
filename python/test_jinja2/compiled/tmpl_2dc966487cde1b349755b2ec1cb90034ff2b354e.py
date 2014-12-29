from __future__ import division
from jinja2.runtime import LoopContext, TemplateReference, Macro, Markup, TemplateRuntimeError, missing, concat, escape, markup_join, unicode_join, to_string, identity, TemplateNotFound
name = '_base.html'

def root(context):
    if 0: yield None
    yield u'<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset="UTF-8">\n\t<title>'
    for event in context.blocks['title'][0](context):
        yield event
    yield u'</title>\n</head>\n<body>\n'
    for event in context.blocks['body'][0](context):
        yield event
    yield u'\n</body>\n</html>'

def block_body(context):
    if 0: yield None
    yield u'\n'
    for event in context.blocks['h1'][0](context):
        yield event
    yield u'\n'
    for event in context.blocks['content'][0](context):
        yield event
    yield u'\n'

def block_content(context):
    if 0: yield None
    yield u'BLOCK content'

def block_h1(context):
    if 0: yield None
    yield u'BLOCK h1'

def block_title(context):
    if 0: yield None
    yield u'BLOCK title'

blocks = {'body': block_body, 'content': block_content, 'h1': block_h1, 'title': block_title}
debug_info = '5=8&8=11&9=18&10=21&9=29&5=33'