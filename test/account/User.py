"User"

from session import session
from cas import login

class User(session.MagicSession):
    "User"
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, account, email):
        "CreateUser"
        params = {'account': account, 'email': email, 'group': [1, 2]}
        val = self.post('/account/user/', params)
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def update(self, user):
        'UpdateUser'
        params = {'email': user['email'], 'name': user['name']}
        val = self.put('/account/user/%d?authToken=%s'%(user['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def updatepassword(self, user, pwd):
        'UpdateUserPassword'
        params = {'password': pwd, 'email': user['email'], 'name': user['name']}
        val = self.put('/account/user/%d?authToken=%s'%(user['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def find(self, user_id):
        'FindUser'
        val = self.get('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['user']
        return None

    def find_all(self):
        'FindAllUser'
        val = self.get('/account/user/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            if len(val['user']) != 2:
                return False
            return True
        return False


    def destroy(self, user_id):
        'DestroyUser'
        val = self.delete('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = User('http://localhost:8888', login_session.authority_token)
        user = app.create('testUser12', 'rangh@test.com')
        if user:
            user_id = user['id']
            if not app.updatepassword(user, '123'):
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
