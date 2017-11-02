"CatalogTest"

import MagicSession
import LoginTest

def join_str(catalog):
    'JoinStr'
    ret = ''
    for v in catalog:
        ret = '%s,%d'%(ret, v)
    return ret

class CatalogTest(MagicSession.MagicSession):
    'CatalogTest'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token


    def create(self, name, description, catalogs):
        'create'
        params = {'catalog-name': name, 'catalog-description': description, 'catalog-parent': catalogs}
        val = self.post('/content/catalog/?authToken=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            print 'create catalog success'
            return val['Catalog']

        print 'create catalog failed'
        return None

    def destroy(self, catalog_id):
        'destroy'
        val = self.delete('/content/catalog/%s?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'destroy catalog success'
            return True

        print 'destroy catalog failed'
        return False

    def update(self, catalog):
        'update'
        catalogs = join_str(catalog['Catalog'])
        params = {'catalog-name': catalog['Name'], 'catalog-description': catalog['Description'], 'catalog-parent': catalogs}
        val = self.put('/content/catalog/%s?authToken=%s'%(catalog['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update catalog success'
            return val['Catalog']

        print 'update catalog failed'
        return None

    def query(self, catalog_id):
        'query'
        val = self.get('/content/catalog/%d?authToken=%s'%(catalog_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'query catalog success'
            return val['Catalog']

        print 'query catalog failed'
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/catalog/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            print 'query_all catalog success'
            return val['Catalog']

        print 'query_all catalog failed'
        return None

if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:
        APP = CatalogTest('http://localhost:8888', LOGIN.authority_token)
        CATALOG = APP.create('testCatalog', 'testDescription', '8,9')
        if CATALOG:
            CATALOG_ID = CATALOG['ID']
            CATALOG['Description'] = 'aaaaaa'
            CATALOG['Catalog'] = [8,9,10]
            CATALOG = APP.update(CATALOG)
            if not CATALOG:
                print 'update catalog failed'
            elif len(CATALOG['Catalog']) != 3:
                print 'update catalog failed, catalog len invalid'
            else:
                pass
            CATALOG = APP.query(CATALOG_ID)
            if not CATALOG:
                print 'query catalog failed'
            elif cmp(CATALOG['Description'],'aaaaaa') != 0:
                print 'update catalog failed, description invalid'

            if len(APP.query_all()) <= 0:
                print 'query_all catalog failed'
            
            APP.destroy(CATALOG_ID)
        else:
            print 'create catalog failed'

        LOGIN.logout(LOGIN.authority_token)
