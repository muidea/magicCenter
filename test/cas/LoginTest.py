'LoginTest'

from session import MagicSession
import sys

class LoginTest(MagicSession.MagicSession):
    'LoginTest'
    def __init__(self, base_url):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = ''
        self.current_user = None

    def login(self, account, password):
        'login'
        print('login')
        params = {'account': account, 'password': password}
        val = self.post('/cas/user/', params)
        if val and val['ErrCode'] == 0:
            self.authority_token = val['AuthToken']
            self.current_user = val['User']
            print('login success')
            return True
        print('login failed')
        return False

    def logout(self, auth_token):
        'logout'
        print('logout')
        val = self.delete('/cas/user/?authToken=%s'%auth_token)
        if val and val['ErrCode'] == 0:
            print('logout success')
        else:
            print('logout failed')

    def status(self, auth_token):
        'status'
        print('status')
        val = self.get('/cas/user/?authToken=%s'%auth_token)
        if val and val['ErrCode'] == 0:
            print('get user status success')
        else:
            print('get user status failed')

def main():
    APP = LoginTest('http://localhost:8888')
    APP.login('rangh@126.com', '123')
    APP.status(APP.authority_token)
    APP.logout('123')
    APP.logout(APP.authority_token)