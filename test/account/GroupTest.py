"GruopTest"

import MagicSession
import LoginTest


class GroupTest(MagicSession.MagicSession):
    "GroupTest"
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, description):
        "CreateGroup"
        params = {'name': name, 'description': description}
        val = self.post('/account/group/?authToken=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            print 'create group success'
            return val['Group']
        else:
            print 'create group failed'
            return None

    def save(self, group):
        "UpdateGroup"
        params = {'name': group['Name'], 'description': group['Description']}
        val = self.put('/account/group/%d?authToken=%s'%(group['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update group success'
            return val['Group']
        else:
            print 'update group failed'
            return None

    def find(self, group_id):
        "FindGroup"
        val = self.get('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'find group success'
            return val['Group']
        else:
            print 'find group failed'
            return None

    def find_all(self):
        "FindAllGroup"
        val = self.get('/account/group/')
        if val and val['ErrCode'] == 0:
            if len(val['Group']) < 0:
                print 'find all group failed'
                return False

            print 'find all group success'
            return True

        return False

    def destroy(self, group_id):
        "DestroyGroup"
        val = self.delete('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'destroy group success'
            return True
        else:
            print 'destroy group failed'
            return False

if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:    
        APP = GroupTest('http://localhost:8888', LOGIN.authority_token)
        GROUP = APP.create('testGorup1', 'test description')
        if GROUP:
            GROUP_ID = GROUP['ID']
            GROUP['Description'] = 'aaaaaa'
            GROUP = APP.save(GROUP)
            if GROUP and cmp(GROUP['Description'], 'aaaaaa') != 0:
                print 'update group failed'

            GROUP = APP.find(GROUP_ID)
            if GROUP:
                print GROUP
                if cmp(GROUP['Description'], 'aaaaaa') != 0:
                    print 'find group failed'
            else:
                print 'find group failed'

            APP.find_all()

            if not APP.destroy(GROUP_ID):
                print 'destroy group failed'

        else:
            print 'create group failed'
        
        LOGIN.logout(LOGIN.authority_token)