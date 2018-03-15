
import common

class Catalog:
    def __init__(self):
        self.name = common.name()
        self.description = common.paragraph()
        self.catalog = []

    def mock(self):
        return self.name, self.description, self.catalog

class Article:
    def __init__(self):
        self.title = ''
        self.content =''
        self.catalog = []

    def mock(self):
        return self.title,self.content,self.catalog



if __name__ == '__main__':
    catalog = Catalog()
    name, desc, catalog = catalog.mock()
    print(name)