"LinkTest"
from session import MagicSession
from cas import LoginTest

class LinkTest(MagicSession.MagicSession):
    'LinkTest'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, url, logo, catalogs):
        'create'
        params = {'name': name, 'url': url, 'logo': logo, 'catalog': str(catalogs)}
        val = self.post('/content/link/?authToken=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            return val['Link']
        return None

    def destroy(self, link_id):
        'destroy'
        val = self.delete('/content/link/%s?authToken=%s'%(link_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return True
        return False

    def update(self, link):
        'update'
        params = {'name': link['Name'], 'url': link['URL'], 'logo': link['Logo'], 'catalog': str(link['Catalog'])}
        val = self.put('/content/link/%s?authToken=%s'%(link['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            return val['Link']
        return None

    def query(self, link_id):
        'query'
        val = self.get('/content/link/%d?authToken=%s'%(link_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return val['Link']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/link/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            return val['Link']
        return None

def main():
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:
        APP = LinkTest('http://localhost:8888', LOGIN.authority_token)
        LINK = APP.create('testLink', 'test link url', 'test link logo', [8,9])
        if LINK:
            LINK_ID = LINK['ID']
            LINK['URL'] = 'aaaaaa, bb dsfsdf  erewre aa'
            LINK['Logo'] = 'test link logo'
            LINK['Catalog'] = [8,9,10]
            LINK = APP.update(LINK)
            if not LINK:
                print('update link failed')
            elif len(LINK['Catalog']) != 3:
                print('update link failed, link len invalid')
            else:
                pass
            LINK = APP.query(LINK_ID)
            if not LINK:
                print('query link failed')
            elif not (LINK['URL'] == 'aaaaaa, bb dsfsdf  erewre aa'):
                print('update link failed, content invalid')

            if len(APP.query_all()) <= 0:
                print('query_all link failed')
            
            APP.destroy(LINK_ID)
        else:
            print('create link failed')

        LOGIN.logout(LOGIN.authority_token)
