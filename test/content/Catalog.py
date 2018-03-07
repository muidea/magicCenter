"Catalog"

from session import MagicSession
from cas import Login

class Catalog(MagicSession.MagicSession):
    'Catalog'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
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
    LOGIN = Login.Login('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:
        APP = Catalog('http://localhost:8888', LOGIN.authority_token)
        CATALOG = APP.create('testCatalog', 'testDescription', [8,9])
        if CATALOG:
            CATALOG_ID = CATALOG['id']
            CATALOG['description'] = 'aaaaaa'
            CATALOG['catalog'] = [8,9,10]
            CATALOG = APP.update(CATALOG)
            if not CATALOG:
                print('update catalog failed')
            elif len(CATALOG['catalog']) != 3:
                print('update catalog failed, catalog len invalid')
            else:
                pass
            CATALOG = APP.query(CATALOG_ID)
            if not CATALOG:
                print('query catalog failed')
            elif not (CATALOG['description'] == 'aaaaaa'):
                print('update catalog failed, description invalid')

            if len(APP.query_all()) <= 0:
                print('query_all catalog failed')
            
            APP.destroy(CATALOG_ID)
        else:
            print('create catalog failed')

        LOGIN.logout(LOGIN.authority_token)
