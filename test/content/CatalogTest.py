"CatalogTest"

from session import MagicSession
from cas import LoginTest

class CatalogTest(MagicSession.MagicSession):
    'CatalogTest'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token


    def create(self, name, description, catalogs):
        'create'
        params = {'name': name, 'description': description, 'catalog': str(catalogs)}
        val = self.post('/content/catalog/?authToken=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            return val['Catalog']
        return None

    def destroy(self, catalog_id):
        'destroy'
        val = self.delete('/content/catalog/%s?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return True
        return False

    def update(self, catalog):
        'update'
        params = {'name': catalog['Name'], 'description': catalog['Description'], 'catalog': str(catalog['Catalog'])}
        val = self.put('/content/catalog/%s?authToken=%s'%(catalog['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            return val['Catalog']
        return None

    def query(self, catalog_id):
        'query'
        val = self.get('/content/catalog/%d?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return val['Catalog']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/catalog/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            return val['Catalog']
        return None

def main():
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:
        APP = CatalogTest('http://localhost:8888', LOGIN.authority_token)
        CATALOG = APP.create('testCatalog', 'testDescription', [8,9])
        if CATALOG:
            CATALOG_ID = CATALOG['ID']
            CATALOG['Description'] = 'aaaaaa'
            CATALOG['Catalog'] = [8,9,10]
            CATALOG = APP.update(CATALOG)
            if not CATALOG:
                print('update catalog failed')
            elif len(CATALOG['Catalog']) != 3:
                print('update catalog failed, catalog len invalid')
            else:
                pass
            CATALOG = APP.query(CATALOG_ID)
            if not CATALOG:
                print('query catalog failed')
            elif not (CATALOG['Description'] == 'aaaaaa'):
                print('update catalog failed, description invalid')

            if len(APP.query_all()) <= 0:
                print('query_all catalog failed')
            
            APP.destroy(CATALOG_ID)
        else:
            print('create catalog failed')

        LOGIN.logout(LOGIN.authority_token)
