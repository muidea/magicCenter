"MediaTest"

import MagicSession

class MediaTest(MagicSession.MagicSession):
    'MediaTest'
    def __init__(self):
        MagicSession.MagicSession.__init__(self)


    def create(self, title, content, catalogs):
        'create'
        pass

    def destroy(self, media_id):
        'destroy'
        pass

    def update(self, media):
        'update'
        pass

    def query(self, media_id):
        'query'
        pass

    def query_all(self):
        'query_all'
        pass