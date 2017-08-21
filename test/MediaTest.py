"MediaTest"

import MagicSession
import LoginTest

def join_str(catalog):
    'JoinStr'
    ret = ''
    for v in catalog:
        ret = '%s,%d'%(ret, v)
    return ret

class MediaTest(MagicSession.MagicSession):
    'MediaTest'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, url, desc, catalogs):
        'create'
        params = {'media-name': name, 'media-url': url, 'media-desc': desc, 'media-catalog': catalogs}
        val = self.post('/content/media/?token=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            print 'create media success'
            return val['Media']

        print 'create media failed'
        return None

    def destroy(self, media_id):
        'destroy'
        val = self.delete('/content/media/%s/?token=%s'%(media_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'destroy media success'
            return True

        print 'destroy media failed'
        return False

    def update(self, media):
        'update'
        catalogs = join_str(media['Catalog'])
        params = {'media-name': media['Name'], 'media-url': media['URL'], 'media-desc': media['Desc'], 'media-catalog': catalogs}
        val = self.put('/content/media/%s/?token=%s'%(media['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update media success'
            return val['Media']

        print 'update media failed'
        return None

    def query(self, media_id):
        'query'
        val = self.get('/content/media/%d/?token=%s'%(media_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'query media success'
            return val['Media']

        print 'query media failed'
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/media/?token=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            print 'query_all media success'
            return val['Media']

        print 'query_all media failed'
        return None

if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888/api/v1')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:
        APP = MediaTest('http://localhost:8888/api/v1', LOGIN.authority_token)
        MEDIA = APP.create('testMedia', 'test media url', 'test media desc', '8,9')
        if MEDIA:
            MEDIA_ID = MEDIA['ID']
            MEDIA['URL'] = 'aaaaaa, bb dsfsdf  erewre'
            MEDIA['Desc'] = 'aaaaaa, bb'
            MEDIA['Catalog'] = [8,9,10]
            MEDIA = APP.update(MEDIA)
            if not MEDIA:
                print 'update media failed'
            elif len(MEDIA['Catalog']) != 3:
                print 'update media failed, media len invalid'
            else:
                pass
            MEDIA = APP.query(MEDIA_ID)
            if not MEDIA:
                print 'query media failed'
            elif cmp(MEDIA['URL'],'aaaaaa, bb dsfsdf  erewre') != 0:
                print 'update media failed, content invalid'

            if len(APP.query_all()) <= 0:
                print 'query_all media failed'
            
            APP.destroy(MEDIA_ID)
        else:
            print 'create media failed'

        LOGIN.logout(LOGIN.authority_token)
