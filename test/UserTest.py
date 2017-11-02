"UserTest"

import MagicSession
import LoginTest

class UserTest(MagicSession.MagicSession):
    "UserTest"
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, account, email):
        "CreateUser"
        params = {'user-account': account, 'user-email': email, 'user-groups': "1,2"}
        val = self.post('/account/user/', params)
        if val and val['ErrCode'] == 0:
            print 'create user success'
            return val['User']

        print 'create user failed'
        print val
        return None

    def update(self, user):
        'UpdateUser'
        params = {'user-email': user['Email'], 'user-name': user['Name']}
        val = self.put('/account/user/%d?authToken=%s'%(user['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update user success'
            return val['User']
        print 'update user failed'
        return None

    def updatepassword(self, user, pwd):
        'UpdateUserPassword'
        params = {'user-password': pwd, 'user-email': user['Email'], 'user-name': user['Name']}
        val = self.put('/account/user/%d?authToken=%s'%(user['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update user password success'
            return val['User']
        print 'update user password failed'
        return None

    def find(self, user_id):
        'FindUser'
        val = self.get('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'find user success'
            return val['User']
        print 'find user failed'
        return None

    def find_all(self):
        'FindAllUser'
        val = self.get('/account/user/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            if len(val['User']) != 2:
                print 'find all user failed'
                return False

            return True

        return False


    def destroy(self, user_id):
        'DestroyUser'
        val = self.delete('/account/user/%d?authToken=%s'%(user_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'destroy user success'
            return True
        print 'destroy user failed'
        return False

if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:    
        APP = UserTest('http://localhost:8888', LOGIN.authority_token)
        USER = APP.create('testUser12', 'rangh@test.com')
        if USER:
            USER_ID = USER['ID']
            print USER
            USER_NEW = APP.updatepassword(USER, '123')
            #if USER_NEW:
                #LOGIN = LoginTest.LoginTest('http://localhost:8888')
                #if not LOGIN.login(USER['Account'], '123'):
                #    print 'update user password failed'
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

            APP.find_all()
        else:
            print 'create user failed'

        LOGIN.logout(LOGIN.authority_token)
