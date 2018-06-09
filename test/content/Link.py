"Link"
from session import session
from cas import cas

class Link:
    'Link'
    def __init__(self, work_session, auth_token):
        self.session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, name, url, logo, catalogs):
        'create'
        params = {'name': name, 'url': url, 'logo': logo, 'catalog': catalogs}
        val = self.session.post('/content/link/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['link']
        return None

    def destroy(self, link_id):
        'destroy'
        val = self.session.delete('/content/link/%s?authToken=%s'%(link_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

    def update(self, link):
        'update'
        params = {'name': link['name'], 'url': link['url'], 'logo': link['logo'], 'catalog': link['catalog']}
        val = self.session.put('/content/link/%s?authToken=%s'%(link['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['link']
        return None

    def query(self, link_id):
        'query'
        val = self.session.get('/content/link/%d?authToken=%s'%(link_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['link']
        return None

    def find_all(self):
        'findAllLink'
        val = self.session.get('/content/link/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['link']
        return None

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    cas_session = cas.Cas(work_session)
    if not cas_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Link(work_session, cas_session.authority_token)
        link = app.create('testLink', 'test link url', 'test link logo', [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}])
        if link:
            link_id = link['id']
            link['url'] = 'aaaaaa, bb dsfsdf  erewre aa'
            link['logo'] = 'test link logo'
            link['catalog'] = [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}, {'id':0, 'name': 'ca3'}]
            link = app.update(link)
            if not link:
                print('update link failed')
            elif len(link['catalog']) != 3:
                print('update link failed, link len invalid')
            else:
                pass
            link = app.query(link_id)
            if not link:
                print('query link failed')
            elif link['url'] != 'aaaaaa, bb dsfsdf  erewre aa':
                print('update link failed, content invalid')

            if not app.find_all():
                print('query_all link failed')

            app.destroy(link_id)
        else:
            print('create link failed')

        cas_session.logout(cas_session.authority_token)
