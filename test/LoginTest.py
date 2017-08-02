'LoginTest'

import MagicSession

class MagicCenter(MagicSession.MagicSession):
    'MagicCenter'
    def __init__(self):
        MagicSession.MagicSession.__init__(self)
        self.authority_token = ''
        self.current_user = None

    def login(self, account, password):
        'login'
        params = {'user-account': account, 'user-password': password}
        val = self.post('http://localhost:8888/cas/user/', params)
        if val and val['ErrCode'] == 0:
            self.authority_token = val['AuthToken']
            self.current_user = val['User']
            print 'login success ok'
            print self.current_user
            print self.authority_token
            return True
        else:
            print 'login failed'
            return False

if __name__ == '__main__':
    APP = MagicCenter()
    APP.login('rangh@126.com', '123')
    