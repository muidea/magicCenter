'content manager'


import random

from content import article
from content import catalog
from content import link
from content import media
from mock import common
from session import session
from cas import login

class Content:
    'Content Manager'
    def __init__(self, work_session, auth_token):
        self.current_article_session = article.Article(work_session, auth_token)
        self.current_catalog_session = catalog.Catalog(work_session, auth_token)
        self.current_link_session = link.Link(work_session, auth_token)
        self.current_media_session = media.Media(work_session, auth_token)
        self.articles = {}
        self.catalogs = {}
        self.links = {}
        self.medias = {}

        self.__refresh_article()
        self.__refresh_catalog()
        self.__refresh_link()
        self.__refresh_media()

    def __refresh_article(self):
        all_article = {}
        articles = self.current_article_session.find_all()
        if articles:
            for _, val in  enumerate(articles):
                all_article[val['id']] = val

        self.articles = all_article

    def __refresh_catalog(self):
        all_catalog = {}
        catalogs = self.current_catalog_session.find_all()
        if catalogs:
            for _, val in enumerate(catalogs):
                all_catalog[val['id']] = val

        self.catalogs = all_catalog

    def __refresh_link(self):
        all_link = {}
        links = self.current_link_session.find_all()
        if links:
            for _, val in enumerate(links):
                all_link[val['id']] = val

        self.links = all_link

    def __refresh_media(self):
        all_media = {}
        medias = self.current_media_session.find_all()
        if medias:
            for _, val in enumerate(medias):
                all_media[val['id']] = val

        self.medias = all_media

    def __mock_catalog(self):
        catalog_num = random.randint(1, 10)
        parent_catalog = common.picker_dict(self.catalogs, catalog_num)
        catalogs = []
        for key, val in parent_catalog.items():
            catalogs.append({'id':key, 'name': val['name']})

        return common.word(), common.paragraph(), catalogs

    def __mock_article(self):
        catalog_num = random.randint(1, 10)
        parent_catalog = common.picker_dict(self.catalogs, catalog_num)
        catalogs = []
        for key, val in parent_catalog.items():
            catalogs.append({'id':key, 'name': val['name']})

        return common.title(), common.content(), catalogs

    def __mock_link(self):
        catalog_num = random.randint(1, 10)
        parent_catalog = common.picker_dict(self.catalogs, catalog_num)
        catalogs = []
        for key, val in parent_catalog.items():
            catalogs.append({'id':key, 'name': val['name']})

        return common.title(), common.url(), common.url(), catalogs

    def __mock_media(self):
        catalog_num = random.randint(1, 10)
        parent_catalog = common.picker_dict(self.catalogs, catalog_num)
        catalogs = []
        for key, val in parent_catalog.items():
            catalogs.append({'id':key, 'name': val['name']})

        return common.title(), common.url(), common.paragraph(), catalogs

    def verify(self):
        'verify account infomation between local and server'
        local_article = self.articles
        self.__refresh_article()
        server_article = self.articles

        local_count = len(local_article)
        server_count = len(server_article)

        if local_count > server_count:
            for k in enumerate(local_article.keys()):
                if k not in server_article:
                    print('cant\'t find article in server_article, id:' + k)
                    break
        elif local_count < server_count:
            for k in enumerate(server_article.keys()):
                if k not in local_article:
                    print('cant\'t find article in local_article, id:' + k)
                    break

        local_catalog = self.catalogs
        self.__refresh_catalog()
        server_catalog = self.catalogs

        local_count = len(local_catalog)
        server_count = len(server_catalog)

        if local_count > server_count:
            for k in enumerate(local_catalog.keys()):
                if k not in server_catalog:
                    print('cant\'t find catalog in server_catalog, id:' + k)
                    break
        elif local_count < server_count:
            for k in enumerate(server_catalog.keys()):
                if k not in local_catalog:
                    print('cant\'t find catalog in local_catalog, id:' + k)
                    break

        local_link = self.links
        self.__refresh_link()
        server_link = self.links

        local_count = len(local_link)
        server_count = len(server_link)

        if local_count > server_count:
            for k in enumerate(local_link.keys()):
                if k not in server_link:
                    print('cant\'t find link in server_link, id:' + k)
                    break
        elif local_count < server_count:
            for k in enumerate(server_link.keys()):
                if k not in local_link:
                    print('cant\'t find link in local_link, id:' + k)
                    break

        local_media = self.medias
        self.__refresh_media()
        server_media = self.medias

        local_count = len(local_media)
        server_count = len(server_media)

        if local_count > server_count:
            for k in enumerate(local_media.keys()):
                if k not in server_media:
                    print('cant\'t find media in server_media, id:' + k)
                    break
        elif local_count < server_count:
            for k in enumerate(server_media.keys()):
                if k not in local_media:
                    print('cant\'t find media in local_media, id:' + k)
                    break


    def refresh(self, auth_token):
        'refresh account info'
        self.current_article_session.refresh_token(auth_token)
        self.current_catalog_session.refresh_token(auth_token)
        self.current_link_session.refresh_token(auth_token)
        self.current_media_session.refresh_token(auth_token)

        self.__refresh_article()
        self.__refresh_catalog()
        self.__refresh_link()
        self.__refresh_media()


    def remove(self):
        'remove'
        article_num = random.randint(1, 10)
        article_dict = common.picker_dict(self.articles, article_num)
        for key in article_dict.keys():
            if key > 0:
                if self.current_article_session.destroy(key):
                    self.articles.pop(key)

        catalog_num = random.randint(1, 10)
        catalog_dict = common.picker_dict(self.catalogs, catalog_num)
        for key in catalog_dict.keys():
            if key > 0:
                if self.current_catalog_session.destroy(key):
                    self.catalogs.pop(key)

        link_num = random.randint(1, 10)
        link_dict = common.picker_dict(self.links, link_num)
        for key in link_dict.keys():
            if key > 0:
                if self.current_link_session.destroy(key):
                    self.links.pop(key)

        media_num = random.randint(1, 10)
        media_dict = common.picker_dict(self.medias, media_num)
        for key in media_dict.keys():
            if key > 0:
                if self.current_media_session.destroy(key):
                    self.medias.pop(key)

    def mock(self):
        'mock'
        catalog_count = random.randint(1, 10)
        while catalog_count > 0:
            catalog_count = catalog_count - 1
            name, desc, catalogs = self.__mock_catalog()
            new_catalog = self.current_catalog_session.create(name, desc, catalogs)
            if new_catalog is None:
                print('create catalog failed')
                break
            else:
                self.catalogs[new_catalog['id']] = new_catalog

        article_count = random.randint(1, 10)
        while article_count > 0:
            article_count = article_count -1
            title, content, catalogs = self.__mock_article()
            new_article = self.current_article_session.create(title, content, catalogs)
            if new_article is None:
                print('create article failed')
                break
            else:
                self.articles[new_article['id']] = new_article

        link_count = random.randint(1, 10)
        while link_count > 0:
            link_count = link_count -1
            name, url, logo, catalogs = self.__mock_link()
            new_link = self.current_link_session.create(name, url, logo, catalogs)
            if new_link is None:
                print('create link failed')
                break
            else:
                self.links[new_link['id']] = new_link

        media_count = random.randint(1, 10)
        while media_count > 0:
            media_count = media_count -1
            name, url, desc, catalogs = self.__mock_media()
            new_media = self.current_media_session.create(name, url, desc, catalogs)
            if new_media is None:
                print('create media failed')
                break
            else:
                self.medias[new_media['id']] = new_media
def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    login_session = login.Login(work_session)
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Content(work_session, login_session.authority_token)
        app.refresh(login_session.authority_token)
        app.mock()
        app.remove()
        app.verify()

        login_session.logout(login_session.authority_token)
