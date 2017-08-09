"ArticleTest"

import MagicSession

class ArticleTest(MagicSession.MagicSession):
    'ArticleTest'
    def __init__(self):
        MagicSession.MagicSession.__init__(self)

    def create(self, title, content, catalogs):
        'create'
        pass

    def destroy(self, article_id):
        'destroy'
        pass

    def update(self, article):
        'update'
        pass

    def query(self, article_id):
        'query'
        pass

    def query_all(self):
        'query_all'
        pass
