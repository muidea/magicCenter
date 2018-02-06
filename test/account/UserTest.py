"UserTest"

from session import MagicSession
from cas import LoginTest

class UserTest(MagicSession.MagicSession):
    "UserTest"
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, account, email):
        "CreateUser"
        params = {'account': account, 'email': email, 'group': str([1,2])}
        val = self.post('/account/user/', params)
        if val and val['ErrCode'] == 0:
            return val['User']
        return None

    def update(self, user):
        'UpdateUser'
        params = {'email': user['Email'], 'name': user['Name']}
        val = self.put('/account/user/%d?authToken=%s'%(user['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            return val['User']
        return None

    def updatepassword(self, user, pwd):
        'UpdateUserPassword'
        params = {'password': pwd, 'email': user['Email'], 'name': user['Name']}
        val = self.put('/account/user/%d?authToken=%s'%(user['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            return val['User']
        return None

    def find(self, user_id):
        'FindUser'
        val = self.get('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return val['User']
        return None

    def find_all(self):
        'FindAllUser'
        val = self.get('/account/user/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            if len(val['User']) != 2:
                return False
            return True
        return False


    def destroy(self, user_id):
        'DestroyUser'
        val = self.delete('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return True
        return False

def main():
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:    
        APP = UserTest('http://localhost:8888', LOGIN.authority_token)
        USER = APP.create('testUser12', 'rangh@test.com')
        if USER:
            USER_ID = USER['ID']
            if not APP.updatepassword(USER, '123'):
                print("updatepassword failed")

            USER['Name'] = '11223'
            USER = APP.update(USER)
            if USER:
                if not (USER['Name'] == '11223'):
                    print('update user failed')
            else:
                print('update user failed')

            USER = APP.find(USER_ID)
            if USER:
                if not (USER['Name'] == '11223'):
                    print('find user failed')
            else:
                print('find user failed')

            APP.destroy(USER_ID)

            APP.find_all()
        else:
            print('create user failed')

        LOGIN.logout(LOGIN.authority_token)
