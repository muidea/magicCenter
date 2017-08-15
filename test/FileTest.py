"FileTest"


import MagicSession
import LoginTest


class FileRegistryTest(MagicSession.MagicSession):
    "FileRegistryTest"
    def __init__(self):
        MagicSession.MagicSession.__init__(self)
    
    def upload(self, file_name):
        'upload file'
        key_name = 'test-name'
        params = {key_name: open(file_name, 'rb')}
        val = self.post('http://localhost:8888/fileregistry/?key-name=%s'%key_name, params)
        if val and val['ErrCode'] == 0:
            print 'upload file success'
            return val['FilePath']

        print 'upload file failed'
        return None

    def delete_file(self, token):
        'delete file'
        pass

    