"MagicSession"

import urllib2
import urllib
import cookielib
import json

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
        ret = None
        try:
            request = urllib2.Request(url, urllib.urlencode(params))
            request.get_method = lambda: 'POST'
            val = self.opener.open(request).readlines()[0]
            ret = json.loads(val)
        except urllib2.URLError:
            print 'post request exception'
        except ValueError:
            print ('decode json exception,val:{1}', val)

        return ret

    def get(self, url):
        "Get"
        ret = None
        try:
            request = urllib2.Request(url)
            request.get_method = lambda: 'GET'
            val = self.opener.open(request).readlines()[0]
            ret = json.loads(val)
        except urllib2.URLError:
            print 'post request exception'
        except ValueError:
            print ('decode json exception,val:{1}', val)

        return ret        

    def put(self, url, params):
        "Put"
        ret = None
        try:
            request = urllib2.Request(url, urllib.urlencode(params))
            request.get_method = lambda: 'PUT'
            val = self.opener.open(request).readlines()[0]
            ret = json.loads(val)
        except urllib2.URLError:
            print 'post request exception'
        except ValueError:
            print ('decode json exception,val:{1}', val)

        return ret

    def delete(self, url):
        "Delete"
        ret = None
        try:
            request = urllib2.Request(url)
            request.get_method = lambda: 'DELETE'
            val = self.opener.open(request).readlines()[0]
            ret = json.loads(val)
        except urllib2.URLError:
            print 'post request exception'
        except ValueError:
            print ('decode json exception,val:{1}', val)

        return ret            
