"Media"

from session import session
from cas import login

class Media(session.MagicSession):
    'Media'
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, url, desc, catalogs):
        'create'
        params = {'name': name, 'url': url, 'description': desc, 'catalog': catalogs}
        val = self.post('/content/media/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['media']
        return None

    def destroy(self, media_id):
        'destroy'
        val = self.delete('/content/media/%s?authToken=%s'%(media_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

    def update(self, media):
        'update'
        params = {'name': media['name'], 'url': media['url'], 'desc': media['description'], 'catalog': media['catalog']}
        val = self.put('/content/media/%s?authToken=%s'%(media['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['media']
        return None

    def query(self, media_id):
        'query'
        val = self.get('/content/media/%d?authToken=%s'%(media_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['media']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/media/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['media']
        return None

def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Media('http://localhost:8888', login_session.authority_token)
        media = app.create('testMedia', 'test media url', 'test media desc', [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}])
        if media:
            media_id = media['id']
            media['url'] = 'aaaaaa, bb dsfsdf  erewre'
            media['description'] = 'aaaaaa, bb'
            media['catalog'] = [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}, {'id':0, 'name': 'ca3'}]
            media = app.update(media)
            if not media:
                print('update media failed')
            elif len(media['catalog']) != 3:
                print('update media failed, media len invalid')
            else:
                pass
            media = app.query(media_id)
            if not media:
                print('query media failed')
            elif media['url'] != 'aaaaaa, bb dsfsdf  erewre':
                print('update media failed, content invalid')

            if len(app.query_all()) <= 0:
                print('query_all media failed')
            app.destroy(media_id)
        else:
            print('create media failed')

        login_session.logout(login_session.authority_token)
