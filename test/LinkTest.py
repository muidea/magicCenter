"LinkTest"

import MagicSession
import LoginTest

def join_str(catalog):
    'JoinStr'
    ret = ''
    for v in catalog:
        ret = '%s,%d'%(ret, v)
    return ret

class LinkTest(MagicSession.MagicSession):
    'LinkTest'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, url, logo, catalogs):
        'create'
        params = {'name': name, 'url': url, 'logo': logo, 'catalog': catalogs}
        val = self.post('/content/link/?authToken=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            print 'create link success'
            return val['Link']

        print 'create link failed'
        return None

    def destroy(self, link_id):
        'destroy'
        val = self.delete('/content/link/%s?authToken=%s'%(link_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'destroy link success'
            return True

        print 'destroy link failed'
        return False

    def update(self, link):
        'update'
        catalogs = join_str(link['Catalog'])
        params = {'name': link['Name'], 'url': link['URL'], 'logo': link['Logo'], 'catalog': catalogs}
        val = self.put('/content/link/%s?authToken=%s'%(link['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update link success'
            return val['Link']

        print 'update link failed'
        return None

    def query(self, link_id):
        'query'
        val = self.get('/content/link/%d?authToken=%s'%(link_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'query link success'
            return val['Link']

        print 'query link failed'
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/link/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            print 'query_all link success'
            return val['Link']

        print 'query_all link failed'
        return None

if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:
        APP = LinkTest('http://localhost:8888', LOGIN.authority_token)
        LINK = APP.create('testLink', 'test link url', 'test link logo', '8,9')
        if LINK:
            LINK_ID = LINK['ID']
            LINK['URL'] = 'aaaaaa, bb dsfsdf  erewre aa'
            LINK['Logo'] = 'test link logo'
            LINK['Catalog'] = [8,9,10]
            LINK = APP.update(LINK)
            if not LINK:
                print 'update link failed'
            elif len(LINK['Catalog']) != 3:
                print 'update link failed, link len invalid'
            else:
                pass
            LINK = APP.query(LINK_ID)
            if not LINK:
                print 'query link failed'
            elif cmp(LINK['URL'],'aaaaaa, bb dsfsdf  erewre aa') != 0:
                print 'update link failed, content invalid'

            if len(APP.query_all()) <= 0:
                print 'query_all link failed'
            
            APP.destroy(LINK_ID)
        else:
            print 'create link failed'

        LOGIN.logout(LOGIN.authority_token)
