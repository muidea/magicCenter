"UserTest"

import MagicSession
import LoginTest

class UserTest(MagicSession.MagicSession):
    "UserTest"
    def __init__(self):
        MagicSession.MagicSession.__init__(self)

    def create(self, account, email):
        "CreateUser"
        params = {'user-account': account, 'user-email': email}
        val = self.post('http://localhost:8888/account/user/', params)
        if val and val['ErrCode'] == 0:
            print 'create user success'
            return val['User']
        else:
            print 'create user failed'
            return None

    def update(self, user):
        'UpdateUser'
        params = {'user-email': user['Email'], 'user-name': user['Name']}
        val = self.put('http://localhost:8888/account/user/%d/'%user['ID'], params)
        if val and val['ErrCode'] == 0:
            print 'update user success'
            return val['User']
        else:
            print 'update user failed'
            return None

    def updatepassword(self, user, pwd):
        'UpdateUserPassword'
        params = {'user-password': pwd, 'user-email': user['Email'], 'user-name': user['Name']}
        val = self.put('http://localhost:8888/account/user/%d/'%user['ID'], params)
        if val and val['ErrCode'] == 0:
            print 'update user password success'
            return val['User']
        else:
            print 'update user password failed'
            return None
        
    def find(self, user_id):
        'FindUser'
        val = self.get('http://localhost:8888/account/user/%d/'%user_id)
        if val and val['ErrCode'] == 0:
            print 'find user success'
            return val['User']
        else:
            print 'find user failed'
            return None

    def destroy(self, user_id):
        'DestroyUser'
        val = self.delete('http://localhost:8888/account/user/%d/'%user_id)
        if val and val['ErrCode'] == 0:
            print 'destroy user success'
            return True
        else:
            print 'destroy user failed'
            return False

if __name__ == '__main__':
    APP = UserTest()
    USER = APP.create('testUser12', 'rangh@test.com')
    if USER:
        USER_ID = USER['ID']
        print USER
        USER_NEW = APP.updatepassword(USER, '123')
        if USER_NEW:
            LOGIN = LoginTest.LoginTest()
            if not LOGIN.login(USER['Account'], '123'):
                print 'update user password failed'
        USER_NEW = None

        USER['Name'] = '11223'
        USER = APP.update(USER)
        if USER:
            print USER
            if cmp(USER['Name'], '11223') != 0:
                print 'update user failed'
        else:
            print 'update user failed'

        USER = APP.find(USER_ID)
        if USER:
            print USER
            if cmp(USER['Name'], '11223') != 0:
                print 'find user failed'
        else:
            print 'find user failed'

        APP.destroy(USER_ID)
    else:
        print 'create user failed'

