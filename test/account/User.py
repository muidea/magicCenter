"User"

from session import MagicSession
from cas import Login

class User(MagicSession.MagicSession):
    "User"
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, account, email):
        "CreateUser"
        params = {'account': account, 'email': email, 'group': [1,2]}
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
    LOGIN = Login.Login('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:    
        APP = User('http://localhost:8888', LOGIN.authority_token)
        USER = APP.create('testUser12', 'rangh@test.com')
        if USER:
            USER_ID = USER['id']
            if not APP.updatepassword(USER, '123'):
                print("updatepassword failed")

            USER['name'] = '11223'
            USER = APP.update(USER)
            if USER:
                if not (USER['name'] == '11223'):
                    print('update user failed')
            else:
                print('update user failed')

            USER = APP.find(USER_ID)
            if USER:
                if not (USER['name'] == '11223'):
                    print('find user failed')
            else:
                print('find user failed')

            APP.destroy(USER_ID)

            APP.find_all()
        else:
            print('create user failed')

        LOGIN.logout(LOGIN.authority_token)
