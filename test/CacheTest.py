"CacheTest"

import MagicSession

class CacheTest(MagicSession.MagicSession):
    'CacheTest'
    def __init__(self):
        MagicSession.MagicSession.__init__(self)
        self.authority_token = ''

    def put_in(self, data):
        'put_in'
        params = {'cache-value': data, 'cache-age': 100}
        val = self.post('http://localhost:8888/cache/', params)
        if val and val['ErrCode'] == 0:
            self.authority_token = val['Token']
            print 'put in success'
            return True

        print 'put in failed'
        return False        

    def fetch_out(self, token):
        'fetch_out'
        val = self.get('http://localhost:8888/cache/%s/'%token)
        if val and val['ErrCode'] == 0:
            print 'fetch out success'
        else:
            print 'fetch out failed'

    def remove(self, token):
        'query'
        val = self.delete('http://localhost:8888/cache/%s/'%token)
        if val and val['ErrCode'] == 0:
            print 'remove cache success'
        else:
            print 'remove cache failed'

if __name__ == '__main__':
    APP = CacheTest()
    if APP.put_in("Test"):
        APP.fetch_out(APP.authority_token)

        APP.remove(APP.authority_token)
