"FileRegistryTest"


import MagicSession
import LoginTest

class FileRegistryTest(MagicSession.MagicSession):
    "FileRegistryTest"
    def __init__(self, base_url, auth_token):
        MagicSession.MagicSession.__init__(self, base_url)
        self.authority_token = auth_token
    
    def upload_file(self, file_name):
        'upload file'
        key_name = "test-name"
        params = {"test-name": open(file_name, 'rb')}
        val = self.upload('/fileregistry/file/?key-name=%s&authToken=%s'%(key_name, self.authority_token), params)
        if val and val['ErrCode'] == 0:
            print 'upload file success'
            return val['AccessToken']

        print 'upload file failed'
        print val
        return None

    def download_file(self, access_token):
        'download file'
        val = self.get('/fileregistry/file/%s'%access_token)
        if val and val['ErrCode'] == 0:
            print 'download file success'
            return val['RedirectURL']

        print 'download file failed'
        return None

    def delete_file(self, access_token):
        'delete file'
        val = self.delete('/fileregistry/file/%s?authToken=%s'%(access_token, self.authority_token))
        if val and val['ErrCode'] == 0:
            print 'delete file success'
            return True

        print 'delete file failed'
        return False

    
if __name__ == '__main__':
    LOGIN = LoginTest.LoginTest('http://localhost:8888')
    if not LOGIN.login('rangh@126.com', '123'):
        print 'login failed'
    else:    
        APP = FileRegistryTest('http://localhost:8888', LOGIN.authority_token)
        INFO = APP.upload_file('./ArticleTest.py')
        if INFO:
            URL = APP.download_file(INFO)
            if not URL:
                print 'download file failed'
            APP.delete_file(INFO)

        LOGIN.logout(LOGIN.authority_token)
        