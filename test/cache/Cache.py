"Cache"

from session import session
from cas import cas

class Cache(session.MagicSession):
    'Cache'
    def __init__(self, base_url, authorityToken):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = authorityToken

    def put_in(self, data):
        'put_in'
        params = {'value': data, 'age': 100}
        val = self.post('/cache/item/?authToken=%s'%(self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['token']

        return None

    def fetch_out(self, token):
        'fetch_out'
        val = self.get('/cache/item/%s?authToken=%s'%(token, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['cache']

        return None

    def remove(self, token):
        'query'
        val = self.delete('/cache/item/%s?authToken=%s'%(token, self.authority_token))
        if val and val['errorCode'] == 0:
            return True

        return False

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    cas_session = cas.Cas(work_session)    
    if not cas_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Cache('http://localhost:8888', cas_session.authority_token)
        token = app.put_in("Test")
        if token:
            if not app.fetch_out(token):
                print("fetch out failed")

            if not app.remove(token):
                print("remove failed")
        else:
            print("put in cache failed")
    cas_session.logout(cas_session.authority_token)
