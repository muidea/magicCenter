'LoginTest'

import urllib2
import urllib
import cookielib
import json

class MagicCenter(object):
    'MagicCenter'
    def __init__(self):
        self.operate = ''
        self.cookjar = cookielib.LWPCookieJar('magicCenter.cookie')
        try:
            self.cookjar.load(ignore_discard=True)
        except cookielib.LoadError:
            pass
        except IOError:
            pass
        self.opener = urllib2.build_opener(urllib2.HTTPCookieProcessor(self.cookjar))
        urllib2.install_opener(self.opener)

    def login(self, account, password):
        'login'
        params = {'user-account': account, 'user-password': password}
        print 'login.....'
        req = urllib2.Request('http://localhost:8888/cas/user/', urllib.urlencode(params))

        self.operate = self.opener.open(req)
        obj = json.loads(self.operate.readlines()[0])
        if obj['ErrCode'] == 0:
            print 'login success ok'
        else:
            print 'login failed'

        self.cookjar.save(ignore_discard=True)

if __name__ == '__main__':
    APP = MagicCenter()
    APP.login('rangh@126.com', '123')
    