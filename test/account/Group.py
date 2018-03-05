"Gruop"
from session import MagicSession
from cas import Login

class GroupDetail(MagicSession.MagicSession):
    "GroupDetail"
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, description):
        "CreateGroup"
        params = {'name': name, 'description': description}
        val = self.post('/account/group/?authToken=%s'%self.authority_token, params)
        if val and val['ErrorCode'] == 0:
            return val['GroupDetail']
        else:
            return None

    def save(self, group):
        "UpdateGroup"
        params = {'name': group['Name'], 'description': group['Description']}
        val = self.put('/account/group/%d?authToken=%s'%(group['ID'], self.authority_token), params)
        if val and val['ErrorCode'] == 0:
            return val['GroupDetail']
        else:
            return None

    def find(self, group_id):
        "FindGroup"
        val = self.get('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['ErrorCode'] == 0:
            return val['GroupDetail']
        else:
            return None

    def find_all(self):
        "FindAllGroup"
        val = self.get('/account/group/')
        if val and val['ErrorCode'] == 0:
            if len(val['GroupDetail']) < 0:
                return False
            return True

        return False

    def destroy(self, group_id):
        "DestroyGroup"
        val = self.delete('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['ErrorCode'] == 0:
            return True
        else:
            return False

def main():
    LOGIN = Login.Login('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:    
        APP = GroupDetail('http://localhost:8888', LOGIN.authority_token)
        GROUP = APP.create('testGorup1', 'test description')
        if GROUP:
            GROUP_ID = GROUP['ID']
            GROUP['Description'] = 'aaaaaa'
            GROUP = APP.save(GROUP)
            if GROUP and not (GROUP['Description'] == 'aaaaaa'):
                print('update group failed')

            GROUP = APP.find(GROUP_ID)
            if GROUP:
                if not (GROUP['Description'] == 'aaaaaa'):
                    print('find group failed')
            else:
                print('find group failed')

            APP.find_all()

            if not APP.destroy(GROUP_ID):
                print('destroy group failed')

        else:
            print('create group failed')
        
        LOGIN.logout(LOGIN.authority_token)