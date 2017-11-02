"ArticleTest"

import MagicSession
import LoginTest

def join_str(catalog):
    'JoinStr'
    ret = ''
    for v in catalog:
        ret = '%s,%d'%(ret, v)
    return ret

class ArticleTest(MagicSession.MagicSession):
    'ArticleTest'
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token


    def create(self, title, content, catalogs):
        'create'
        params = {'article-title': title, 'article-content': content, 'article-catalog': catalogs}
        val = self.post('/content/article/?authToken=%s'%(self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'create article success'
            return val['Article']

        print 'create article failed'
        return None

    def destroy(self, article_id):
        'destroy'
        val = self.delete('/content/article/%s?authToken=%s'%(article_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'destroy article success'
            return True

        print 'destroy article failed'
        return False

    def update(self, article):
        'update'
        catalogs = join_str(article['Catalog'])
        params = {'article-title': article['Name'], 'article-content': article['Content'], 'article-catalog': catalogs}
        val = self.put('/content/article/%s?authToken=%s'%(article['ID'], self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'update article success'
            return val['Article']

        print 'update article failed'
        return None

    def query(self, article_id):
        'query'
        val = self.get('/content/article/%d?authToken=%s'%(article_id, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'query article success'
            return val['Article']

        print 'query article failed'
        return None

    def query_all(self):
        'query_all'
        val = self.get('/content/article/?authToken=%s'%self.authority_token)
        if val and val['ErrCode'] == 0:
            print 'query_all article success'
            return val['Article']

        print 'query_all article failed'
        return None

if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:
        APP = ArticleTest('http://localhost:8888', LOGIN.authority_token)
        ARTICLE = APP.create('testArticle', 'test article content', '8,9')
        if ARTICLE:
            ARTICLE_ID = ARTICLE['ID']
            ARTICLE['Content'] = 'aaaaaa, bb dsfsdf  erewre'
            ARTICLE['Catalog'] = [8,9,10]
            ARTICLE = APP.update(ARTICLE)
            if not ARTICLE:
                print 'update article failed'
            elif len(ARTICLE['Catalog']) != 3:
                print 'update article failed, article len invalid'
            else:
                pass
            ARTICLE = APP.query(ARTICLE_ID)
            if not ARTICLE:
                print 'query article failed'
            elif cmp(ARTICLE['Content'],'aaaaaa, bb dsfsdf  erewre') != 0:
                print 'update article failed, content invalid'

            if len(APP.query_all()) <= 0:
                print 'query_all article failed'
             
            APP.destroy(ARTICLE_ID)
        else:
            print 'create article failed'

        LOGIN.logout(LOGIN.authority_token)
