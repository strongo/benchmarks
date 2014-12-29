import sys
from jinja2 import TemplateNotFound, ModuleLoader
from jinja2.utils import internalcode


class KeepModuleLoader(ModuleLoader):
    """This loader loads templates from precompiled templates and keep them in sys.modules.
    http://stackoverflow.com/questions/13485294/compiled-templates-with-macros-do-not-work-on-app-engine
    """

    def __init__(self, path):
        super(KeepModuleLoader, self).__init__(path)
        sys.modules[self.module.__name__] = self.module # Replace weak reference with hard one.

    @internalcode
    def load(self, environment, name, globals=None):
        key = self.get_template_key(name)
        module = '%s.%s' % (self.package_name, key)
        mod = getattr(self.module, module, None)
        if mod is None:
            try:
                mod = __import__(module, None, None, ['root'])
            except ImportError:
                raise TemplateNotFound(name)

            # remove the entry from sys.modules, we only want the attribute
            # on the module object we have stored on the loader.
            # sys.modules.pop(module, None) - Commented as we want to keep module in sys.modules

        return environment.template_class.from_module_dict(
            environment, mod.__dict__, globals)