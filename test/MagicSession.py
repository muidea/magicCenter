"MagicSession"

import urllib2
import urllib
import cookielib

class MagicSession(object):
    "MagicSession"
    def __init__(self):
        print "MagicSession construct"
        self.cookjar = cookielib.LWPCookieJar('magicCenter.cookie')
        try:
            self.cookjar.load(ignore_discard=True)
        except IOError:
            pass
        self.opener = urllib2.build_opener(urllib2.HTTPCookieProcessor(self.cookjar))
        urllib2.install_opener(self.opener)

    def __del__(self):
        print "MagicSession destruct"
        self.cookjar.save(ignore_discard=True)

    def post(self, url, params):
        "Post"
        request = urllib2.Request(url, urllib.urlencode(params))
        request.get_method = lambda: 'POST'
        return self.opener.open(request).readlines()[0]

    def get(self, url):
        "Get"
        request = urllib2.Request(url)
        request.get_method = lambda: 'GET'
        return self.opener.open(request).readlines()[0]

    def put(self, url, params):
        "Put"
        request = urllib2.Request(url, urllib.urlencode(params))
        request.get_method = lambda: 'PUT'
        return self.opener.open(request).readlines()[0]

    def delete(self, url):
        "Delete"
        request = urllib2.Request(url)
        request.get_method = lambda: 'DELETE'
        return self.opener.open(request).readlines()[0]
