"UserTest"

import MagicSession

class User(MagicSession.MagicSession):
    "User"
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

if __name__ == '__main__':
    APP = User()
    USER = APP.create('testUser12','rangh@test.com')
    if USER:
        print USER

