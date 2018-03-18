"Catalog"

from session import session
from cas import login

class Catalog(session.MagicSession):
    'Catalog'
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token


    def create(self, name, description, catalogs):
        'create'
        params = {'name': name, 'description': description, 'catalog': catalogs}
        val = self.post('/content/catalog/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

    def destroy(self, catalog_id):
        'destroy'
        val = self.delete('/content/catalog/%s?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

    def update(self, catalog):
        'update'
        params = {'name': catalog['name'], 'description': catalog['description'], 'catalog': catalog['catalog']}
        val = self.put('/content/catalog/%s?authToken=%s'%(catalog['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

    def query(self, catalog_id):
        'query'
        val = self.get('/content/catalog/%d?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/catalog/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['catalog']
        return None

def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Catalog('http://localhost:8888', login_session.authority_token)
        catalog = app.create('testCatalog', 'testDescription', [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}])
        if catalog:
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

            if len(app.query_all()) <= 0:
                print('query_all catalog failed')

            app.destroy(catalog_id)
        else:
            print('create catalog failed')

        login_session.logout(login_session.authority_token)
