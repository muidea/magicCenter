"LinkTest"

import MagicSession

class LinkTest(MagicSession.MagicSession):
    'LinkTest'
    def __init__(self):
        MagicSession.MagicSession.__init__(self)


    def create(self, title, content, catalogs):
        'create'
        pass

    def destroy(self, link_id):
        'destroy'
        pass

    def update(self, link):
        'update'
        pass

    def query(self, link_id):
        'query'
        pass

    def query_all(self):
        'query_all'
        pass