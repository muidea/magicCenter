'Cas'

from session import session

class Cas:
    'Cas'
    def __init__(self, work_session):
        self.session = work_session
        self.authority_token = ''
        self.current_user = None

    def login(self, account, password):
        'login'
        params = {'account': account, 'password': password}
        val = self.session.post('/cas/user/', params)
        if val and val['errorCode'] == 0:
            self.authority_token = val['onlineEntry']['authToken']
            self.current_user = val['onlineEntry']
            return True
        return False

    def logout(self, auth_token):
        'logout'
        val = self.session.delete('/cas/user/?authToken=%s'%auth_token)
        if val and val['errorCode'] == 0:
            return True
        return False

    def status(self, auth_token):
        'status'
        val = self.session.get('/cas/user/?authToken=%s'%auth_token)
        if val and val['errorCode'] == 0:
            return val['onlineEntry']
        return None

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    app = Cas(work_session)
    app.login('admin@muidea.com', '123')
    app.status(app.authority_token)
    app.logout('123')
    app.logout(app.authority_token)
