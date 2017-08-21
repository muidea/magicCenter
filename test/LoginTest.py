'LoginTest'

import MagicSession

class LoginTest(MagicSession.MagicSession):
    'LoginTest'
    def __init__(self, base_url):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = ''
        self.current_user = None

    def login(self, account, password):
        'login'
        params = {'user_account': account, 'user_password': password}
        val = self.post('/cas/user/', params)
        if val and val['ErrCode'] == 0:
            self.authority_token = val['AuthToken']
            self.current_user = val['User']
            print 'login success'
            return True
        print 'login failed'
        return False

    def logout(self, auth_token):
        'logout'
        val = self.delete('/cas/user/?authToken=%s'%auth_token)
        if val and val['ErrCode'] == 0:
            print 'logout success'
        else:
            print 'logout failed'

if __name__ == '__main__':
    APP = LoginTest('http://localhost:8888/api/v1')
    APP.login('rangh@126.com', '123')
    APP.logout('123')
    APP.logout(APP.authority_token)
