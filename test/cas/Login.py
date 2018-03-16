'Login'

from session import session

class Login(session.MagicSession):
    'Login'
    def __init__(self, base_url):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = ''
        self.current_user = None

    def login(self, account, password):
        'login'
        params = {'account': account, 'password': password}
        val = self.post('/cas/user/', params)
        if val and val['errorCode'] == 0:
            self.authority_token = val['onlineUser']['authToken']
            self.current_user = val['onlineUser']
            return True
        return False

    def logout(self, auth_token):
        'logout'
        val = self.delete('/cas/user/?authToken=%s'%auth_token)
        if val and val['errorCode'] == 0:
            return True
        return False

    def status(self, auth_token):
        'status'
        val = self.get('/cas/user/?authToken=%s'%auth_token)
        if val and val['errorCode'] == 0:
            return val['onlineUser']
        return None

def main():
    'main'
    app = Login('http://localhost:8888')
    app.login('admin@muidea.com', '123')
    app.status(app.authority_token)
    app.logout('123')
    app.logout(app.authority_token)
