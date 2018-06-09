"User"

from session import session
from cas import cas

class User:
    "User"
    def __init__(self, work_session, auth_token):
        self.current_session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, account, password, email, group):
        "CreateUser"
        params = {'account': account, 'password': password, 'email': email, 'group': group}
        val = self.current_session.post('/account/user/', params)
        if val and val['errorCode'] == 0:
            return val['user']

        print('param:')
        print(params)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def update(self, user):
        'UpdateUser'
        params = {'email': user['email'], 'group': user['group']}
        val = self.current_session.put('/account/user/%d?authToken=%s'%(user['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['user']

        print('param:')
        print(params)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def update_password(self, user, pwd):
        'UpdateUserPassword'
        params = {'password': pwd}
        val = self.current_session.put('/account/user/%d?authToken=%s&action=change_password'%(user['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['user']

        print('param:')
        print(params)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def find(self, user_id):
        'FindUser'
        val = self.current_session.get('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['user']

        print('param:')
        print(user_id)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def find_all(self):
        'FindAllUser'
        val = self.current_session.get('/account/user/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['user']

        print('result:')
        print(val)
        print('-----------------------------------------')
        return None


    def destroy(self, user_id):
        'DestroyUser'
        val = self.current_session.delete('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True

        print('param:')
        print(user_id)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return False

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    cas_session = cas.Cas(work_session)
    if not cas_session.login('admin@muidea.com', '123'):
        print('cas failed')
    else:
        app = User(work_session, cas_session.authority_token)
        user = app.create('testUser12', '123', 'rangh@test.com', [{'id': 0, 'name':'系统管理组'}])
        if user:
            user_id = user['id']
            if not app.update_password(user, '123456'):
                print("updatepassword failed")

            user['email'] = '11223@ttt.com'
            user = app.update(user)
            if user:
                if user['email'] != '11223@ttt.com':
                    print('update user failed')
            else:
                print('update user failed')

            user = app.find(user_id)
            if user:
                if user['email'] != '11223@ttt.com':
                    print('find user failed')
            else:
                print('find user failed')

            app.destroy(user_id)

            app.find_all()
        else:
            print('create user failed')

        cas_session.logout(cas_session.authority_token)
