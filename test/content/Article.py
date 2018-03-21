"Article"

from session import session
from cas import login

class Article:
    'Article'
    def __init__(self, work_session, auth_token):
        self.session = work_session
        self.authority_token = auth_token

    def refresh_token(self, auth_token):
        'refreshToken'
        self.authority_token = auth_token

    def create(self, title, content, catalogs):
        'create article'
        params = {'name': title, 'content': content, 'catalog': catalogs}
        val = self.session.post('/content/article/?authToken=%s'%(self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['article']
        return None

    def destroy(self, article_id):
        'destroy'
        val = self.session.delete('/content/article/%s?authToken=%s'%(article_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return True

        print('destroy article failed')
        return False

    def update(self, article):
        'update'
        params = {'title': article['name'], 'content': article['content'], 'catalog': article['catalog']}
        val = self.session.put('/content/article/%s?authToken=%s'%(article['id'], self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['article']
        return None

    def query(self, article_id):
        'query'
        val = self.session.get('/content/article/%d?authToken=%s'%(article_id, self.authority_token))
        if val and val['errorCode'] == 0:
            return val['article']
        return None

    def find_all(self):
        'findAllArticle'
        val = self.session.get('/content/article/?authToken=%s'%self.authority_token)
        if val and val['errorCode'] == 0:
            return val['article']
        return None

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    login_session = login.Login(work_session)
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Article(work_session, login_session.authority_token)
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

            if len(app.find_all()) <= 0:
                print('query_all article failed')

            app.destroy(article_id)
        else:
            print('create article failed')

        login_session.logout(login_session.authority_token)
