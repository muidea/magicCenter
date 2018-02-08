"Media"

from session import MagicSession
from cas import Login

class Media(MagicSession.MagicSession):
    'Media'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, url, desc, catalogs):
        'create'
        params = {'name': name, 'url': url, 'desc': desc, 'catalog': str(catalogs)}
        val = self.post('/content/media/?authToken=%s'%self.authority_token, params)
        if val and val['ErrCode'] == 0:
            return val['Media']
        return None

    def destroy(self, media_id):
        'destroy'
        val = self.delete('/content/media/%s?authToken=%s'%(media_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return True
        return False

    def update(self, media):
        'update'
        params = {'name': media['Name'], 'url': media['URL'], 'desc': media['Desc'], 'catalog': str(media['Catalog'])}
        val = self.put('/content/media/%s?authToken=%s'%(media['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            return val['Media']
        return None

    def query(self, media_id):
        'query'
        val = self.get('/content/media/%d?authToken=%s'%(media_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            return val['Media']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/media/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            return val['Media']
        return None

def main():
    LOGIN = Login.Login('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print('login failed')
    else:
        APP = Media('http://localhost:8888', LOGIN.authority_token)
        MEDIA = APP.create('testMedia', 'test media url', 'test media desc', [8,9])
        if MEDIA:
            MEDIA_ID = MEDIA['ID']
            MEDIA['URL'] = 'aaaaaa, bb dsfsdf  erewre'
            MEDIA['Desc'] = 'aaaaaa, bb'
            MEDIA['Catalog'] = [8,9,10]
            MEDIA = APP.update(MEDIA)
            if not MEDIA:
                print('update media failed')
            elif len(MEDIA['Catalog']) != 3:
                print('update media failed, media len invalid')
            else:
                pass
            MEDIA = APP.query(MEDIA_ID)
            if not MEDIA:
                print('query media failed')
            elif not(MEDIA['URL'] =='aaaaaa, bb dsfsdf  erewre'):
                print('update media failed, content invalid')

            if len(APP.query_all()) <= 0:
                print('query_all media failed')            
            APP.destroy(MEDIA_ID)
        else:
            print('create media failed')

        LOGIN.logout(LOGIN.authority_token)
