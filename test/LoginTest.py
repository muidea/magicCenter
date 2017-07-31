'LoginTest'

import urllib2
import urllib
import cookielib

class MagicCenter(object):
    'MagicCenter'
    def __init__(self):
        self.operate = ''
        self.cookjar = cookielib.LWPCookieJar()
        try:
            self.cookjar.revert('magicCenter.cookie')
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
        data = self.operate.readlines()
        print data
        self.cookjar.save('magicCenter.cookie')

if __name__ == '__main__':
    TT = MagicCenter()
    TT.login('rangh@126.com', '123')
    