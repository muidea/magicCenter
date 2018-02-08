"Cache"

from session import MagicSession
from cas import Login

class Cache(MagicSession.MagicSession):
    'Cache'
    def __init__(self, base_url, authorityToken):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = authorityToken

    def put_in(self, data):
        'put_in'
        params = {'value': data, 'age': 100}
        val = self.post('/cache/item/?authToken=%s'%(self.authority_token), params)
        if val and val['ErrCode'] == 0:
            return val['Token']
        return None

    def fetch_out(self, token):
        'fetch_out'
        val = self.get('/cache/item/%s?authToken=%s'%(token, self.authority_token))
        if val and val['ErrCode'] == 0:
            return val['Cache']
        else:
            return None

    def remove(self, token):
        'query'
        val = self.delete('/cache/item/%s?authToken=%s'%(token, self.authority_token))
        if val and val['ErrCode'] == 0:
            return True
        else:
            return False

def main():
    LOGIN = Login.Login('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:    
        APP = Cache('http://localhost:8888', LOGIN.authority_token)
        token = APP.put_in("Test")
        if token:
            if not APP.fetch_out(token):
                print("fetch out failed")

            if not APP.remove(token):
                print("remove failed")
        else:
            print("put in cache failed")
    LOGIN.logout(LOGIN.authority_token)
