'group.py'

from session import session
from cas import login

class Group:
    "Group"
    def __init__(self, work_session, auth_token):
        self.session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, name, description, catalog):
        "CreateGroup"
        params = {'name': name, 'description': description, 'catalog': catalog}
        val = self.session.post('/account/group/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['group']

        print('param:')
        print(params)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def save(self, group):
        "UpdateGroup"
        params = {'name': group['name'], 'description': group['description']}
        val = self.session.put('/account/group/%d?authToken=%s'%(group['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['group']

        print('param:')
        print(params)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def find(self, group_id):
        "FindGroup"
        val = self.session.get('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['group']

        print('param:')
        print(group_id)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def find_all(self):
        "FindAllGroup"
        val = self.session.get('/account/group/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['group']

        print('result:')
        print(val)
        print('-----------------------------------------')
        return None

    def destroy(self, group_id):
        "DestroyGroup"
        val = self.session.delete('/account/group/%d?authToken=%s'%(group_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True

        print('param:')
        print(group_id)
        print('result:')
        print(val)
        print('-----------------------------------------')
        return False

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    login_session = login.Login(work_session)
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Group(work_session, login_session.authority_token)
        group = app.create('testGorup1', 'test description', {'id':0, 'name':'test'})
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
