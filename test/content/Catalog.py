"Catalog"

from session import session
from cas import login

class Catalog:
    'Catalog'
    def __init__(self, work_session, auth_token):
        self.session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, name, description, catalogs):
        'create'
        params = {'name': name, 'description': description, 'catalog': catalogs}
        val = self.session.post('/content/catalog/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

    def destroy(self, catalog_id):
        'destroy'
        val = self.session.delete('/content/catalog/%s?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

    def update(self, catalog):
        'update'
        params = {'name': catalog['name'], 'description': catalog['description'], 'catalog': catalog['catalog']}
        val = self.session.put('/content/catalog/%s?authToken=%s'%(catalog['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

    def query(self, catalog_id):
        'query'
        val = self.session.get('/content/catalog/%d?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

    def find_all(self):
        'findAllCatalog'
        val = self.session.get('/content/catalog/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    login_session = login.Login(work_session)
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Catalog(work_session, login_session.authority_token)
        catalog = app.create('testCatalog', 'testDescription', [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}])
        if catalog:
            temp = app.create('testCatalog2', 'testDescription', [{'id':catalog['id'], 'name': catalog['name']}, {'id':0, 'name':'ca4'}])
            app.destroy(temp['id'])

            catalog_id = catalog['id']
            catalog['description'] = 'aaaaaa'
            catalog['catalog'] = [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}, {'id':0, 'name': 'ca3'}]
            catalog = app.update(catalog)
            if not catalog:
                print('update catalog failed')
            elif len(catalog['catalog']) != 3:
                print('update catalog failed, catalog len invalid')
            else:
                pass
            catalog = app.query(catalog_id)
            if not catalog:
                print('query catalog failed')
            elif catalog['description'] != 'aaaaaa':
                print('update catalog failed, description invalid')

            if len(app.find_all()) <= 0:
                print('query_all catalog failed')

            app.destroy(catalog_id)
        else:
            print('create catalog failed')

        login_session.logout(login_session.authority_token)
