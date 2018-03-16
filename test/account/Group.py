'group.py'

from session import session
from cas import login

class Group(session.MagicSession):
    "Group"
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, description):
        "CreateGroup"
        params = {'name': name, 'description': description, 'catalog': 0}
        val = self.post('/account/group/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['group']

        return None

    def save(self, group):
        "UpdateGroup"
        params = {'name': group['name'], 'description': group['description']}
        val = self.put('/account/group/%d?authToken=%s'%(group['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['group']

        return None

    def find(self, group_id):
        "FindGroup"
        val = self.get('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['group']

        return None

    def find_all(self):
        "FindAllGroup"
        val = self.get('/account/group/')
        if val and val['errorCode'] == 0:
            if len(val['group']) < 0:
                return False
            return True

        return False

    def destroy(self, group_id):
        "DestroyGroup"
        val = self.delete('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True

        return False

def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Group('http://localhost:8888', login_session.authority_token)
        group = app.create('testGorup1', 'test description')
        if group:
            group_id = group['id']
            group['description'] = 'aaaaaa'
            group = app.save(group)
            if group and (group['description'] != 'aaaaaa'):
                print('update group failed')

            group = app.find(group_id)
            if group:
                if group['description'] != 'aaaaaa':
                    print('find group failed')
            else:
                print('find group failed')

            app.find_all()

            if not app.destroy(group_id):
                print('destroy group failed')

        else:
            print('create group failed')

        login_session.logout(login_session.authority_token)
