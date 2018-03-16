'content'
import common

class Catalog:
    'Catalog'
    def __init__(self):
        self.name = common.name()
        self.description = common.paragraph()
        self.catalog = []

    def mock(self):
        'mock'
        return self.name, self.description, self.catalog

class Article:
    'Article'
    def __init__(self, catalog):
        self.title = common.title()
        self.content = common.content()
        self.catalog = catalog

    def mock(self):
        'mock'
        return self.title, self.content, self.catalog

if __name__ == '__main__':
    current = Catalog()
    name, descr, catalogs = current.mock()
    print(name)
