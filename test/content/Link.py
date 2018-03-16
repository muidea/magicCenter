"Link"
from session import session
from cas import login

class Link(session.MagicSession):
    'Link'
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, name, url, logo, catalogs):
        'create'
        params = {'name': name, 'url': url, 'logo': logo, 'catalog': catalogs}
        val = self.post('/content/link/?authToken=%s'%self.authority_token, params)
        if val and val['errorCode'] == 0:
            return val['link']
        return None

    def destroy(self, link_id):
        'destroy'
        val = self.delete('/content/link/%s?authToken=%s'%(link_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False

    def update(self, link):
        'update'
        params = {'name': link['name'], 'url': link['url'], 'logo': link['logo'], 'catalog': link['catalog']}
        val = self.put('/content/link/%s?authToken=%s'%(link['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['link']
        return None

    def query(self, link_id):
        'query'
        val = self.get('/content/link/%d?authToken=%s'%(link_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['link']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/link/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['link']
        return None

def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Link('http://localhost:8888', login_session.authority_token)
        link = app.create('testLink', 'test link url', 'test link logo', [8, 9])
        if link:
            link_id = link['id']
            link['url'] = 'aaaaaa, bb dsfsdf  erewre aa'
            link['logo'] = 'test link logo'
            link['catalog'] = [8, 9, 10]
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

            if len(app.query_all()) <= 0:
                print('query_all link failed')

            app.destroy(link_id)
        else:
            print('create link failed')

        login_session.logout(login_session.authority_token)
