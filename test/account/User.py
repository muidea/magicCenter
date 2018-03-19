"User"

from session import session
from cas import login

class User:
    "User"
    def __init__(self, work_session, auth_token):
        self.current_session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, account, email, group):
        "CreateUser"
        params = {'account': account, 'email': email, 'group': group}
        val = self.current_session.post('/account/user/', params)
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def update(self, user):
        'UpdateUser'
        params = {'email': user['email'], 'name': user['name']}
        val = self.current_session.put('/account/user/%d?authToken=%s'%(user['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def update_password(self, user, pwd):
        'UpdateUserPassword'
        params = {'password': pwd, 'email': user['email'], 'name': user['name']}
        val = self.current_session.put('/account/user/%d?authToken=%s'%(user['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def find(self, user_id):
        'FindUser'
        val = self.current_session.get('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def find_all(self):
        'FindAllUser'
        val = self.current_session.get('/account/user/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['user']
        return None


    def destroy(self, user_id):
        'DestroyUser'
        val = self.current_session.delete('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    login_session = login.Login(work_session)
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = User(work_session, login_session.authority_token)
        user = app.create('testUser12', 'rangh@test.com', [1, 2])
        if user:
            user_id = user['id']
            if not app.update_password(user, '123'):
                print("updatepassword failed")

            user['name'] = '11223'
            user = app.update(user)
            if user:
                if user['name'] != '11223':
                    print('update user failed')
            else:
                print('update user failed')

            user = app.find(user_id)
            if user:
                if user['name'] != '11223':
                    print('find user failed')
            else:
                print('find user failed')

            app.destroy(user_id)

            app.find_all()
        else:
            print('create user failed')

        login_session.logout(login_session.authority_token)
