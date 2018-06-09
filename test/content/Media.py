"Media"

from session import session
from cas import cas

class Media(session.MagicSession):
    'Media'
    def __init__(self, work_session, auth_token):
        self.session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, name, url, desc, catalogs):
        'create'
        params = {'name': name, 'url': url, 'description': desc, 'catalog': catalogs}
        val = self.session.post('/content/media/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['media']
        return None

    def destroy(self, media_id):
        'destroy'
        val = self.session.delete('/content/media/%s?authToken=%s'%(media_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

    def update(self, media):
        'update'
        params = {'name': media['name'], 'url': media['url'], 'desc': media['description'], 'catalog': media['catalog']}
        val = self.session.put('/content/media/%s?authToken=%s'%(media['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['media']
        return None

    def query(self, media_id):
        'query'
        val = self.session.get('/content/media/%d?authToken=%s'%(media_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['media']
        return None

    def find_all(self):
        'findAllMedia'
        val = self.session.get('/content/media/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['media']
        return None

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    cas_session = cas.Cas(work_session)
    if not cas_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Media(work_session, cas_session.authority_token)
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

            if not app.find_all():
                print('query_all media failed')
            app.destroy(media_id)
        else:
            print('create media failed')

        cas_session.logout(cas_session.authority_token)
