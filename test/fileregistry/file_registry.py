'file_registry'

from session import session
from cas import login

class FileRegistry(session.MagicSession):
    "FileRegistry"
    def __init__(self, base_url, auth_token):
        session.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token

    def upload_file(self, file_name):
        'upload file'
        key_name = "name"
        params = {"name": open(file_name, 'rb')}
        val = self.upload('/fileregistry/file/?key-name=%s&authToken=%s'%(key_name, self.authority_token), params)
        if val and val['errorCode'] == 0:
            return val['accessToken']
        return None

    def download_file(self, access_token):
        'download file'
        val = self.get('/fileregistry/file/%s'%access_token)
        if val and val['errorCode'] == 0:
            return val['redirectURL']
        return None

    def delete_file(self, access_token):
        'delete file'
        val = self.delete('/fileregistry/file/%s?authToken=%s'%(access_token, self.authority_token))
        if val and val['errorCode'] == 0:
            return True
        return False


def main():
    'main'
    login_session = login.Login('http://localhost:8888')
    if not login_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = FileRegistry('http://localhost:8888', login_session.authority_token)
        info = app.upload_file('./ArticleTest.py')
        if info:
            url = app.download_file(info)
            if not url:
                print('download file failed')
            app.delete_file(info)

        login_session.logout(login_session.authority_token)
