"Article"

from session import session
from cas import login

class Article(session.MagicSession):
    'Article'
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def create(self, title, content, catalogs):
        'create article'
        params = {'title': title, 'content': content, 'catalog': catalogs}
        val = self.post('/content/article/?authToken=%s'%(self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['article']
        return None

    def destroy(self, article_id):
        'destroy'
        val = self.delete('/content/article/%s?authToken=%s'%(article_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True

        print('destroy article failed')
        return False

    def update(self, article):
        'update'
        params = {'title': article['name'], 'content': article['content'], 'catalog': article['catalog']}
        val = self.put('/content/article/%s?authToken=%s'%(article['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['article']
        return None

    def query(self, article_id):
        'query'
        val = self.get('/content/article/%d?authToken=%s'%(article_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['article']
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/article/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['article']
        return None

def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Article('http://localhost:8888', login_session.authority_token)
        article = app.create('testArticle', 'test article content', [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}])
        if article:
            article_id = article['id']
            article['content'] = 'aaaaaa, bb dsfsdf  erewre'
            article['catalog'] = [{'id':0, 'name': "ca1"}, {'id':0, 'name':'ca2'}, {'id':0, 'name': 'ca3'}]
            article = app.update(article)
            if not article:
                print('update article failed')
            elif len(article['catalog']) != 3:
                print('update article failed, article len invalid')
            else:
                pass
            article = app.query(article_id)
            if not article:
                print('query article failed')
            elif article['content'] != 'aaaaaa, bb dsfsdf  erewre':
                print('update article failed, content invalid')

            if len(app.query_all()) <= 0:
                print('query_all article failed')

            app.destroy(article_id)
        else:
            print('create article failed')

        login_session.logout(login_session.authority_token)
