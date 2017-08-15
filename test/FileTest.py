"FileTest"


import MagicSession

class FileRegistryTest(MagicSession.MagicSession):
    "FileRegistryTest"
    def __init__(self):
        MagicSession.MagicSession.__init__(self)
    
    def upload_file(self, file_name):
        'upload file'
        key_name = "test-name"
        params = {"test-name": open(file_name, 'rb')}
        val = self.upload('http://localhost:8888/fileregistry/?key-name=%s'%key_name, params)
        if val and val['ErrCode'] == 0:
            print 'upload file success'
            return val['AccessToken']

        print 'upload file failed'
        print val
        return None

    def download_file(self, access_token):
        'download file'
        val = self.get('http://localhost:8888/fileregistry/%s/'%access_token)
        if val and val['ErrCode'] == 0:
            print 'download file success'
            return val['RedirectURL']

        print 'download file failed'
        return None

    def delete_file(self, access_token):
        'delete file'
        val = self.delete('http://localhost:8888/fileregistry/%s/'%access_token)
        if val and val['ErrCode'] == 0:
            print 'delete file success'
            return True

        print 'delete file failed'
        return False

    
if __name__ == '__main__':
    APP = FileRegistryTest()
    INFO = APP.upload_file('./ArticleTest.py')
    if INFO:
        URL = APP.download_file(INFO)
        if not URL:
            print 'download file failed'
        APP.delete_file(INFO)