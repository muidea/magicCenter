'LoginTest'

import json
import MagicSession

class MagicCenter(MagicSession.MagicSession):
    'MagicCenter'
    def __init__(self):
        MagicSession.MagicSession.__init__(self)

    def login(self, account, password):
        'login'

        params = {'user-account': account, 'user-password': password}
        val = self.post('http://localhost:8888/cas/user/', params)
        obj = json.loads(val)
        if obj['ErrCode'] == 0:
            print 'login success ok'
        else:
            print 'login failed'

if __name__ == '__main__':
    APP = MagicCenter()
    APP.login('rangh@126.com', '123')
    