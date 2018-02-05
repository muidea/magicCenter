"MagicSession"

import requests
import json
class MagicSession(object):
    "MagicSession"
    def __init__(self, base_url):
        self.currentSesion = requests.Session()
        self.base_url = base_url

    def post(self, url, params):
        "Post"
        ret = None
        try:
            response = self.currentSesion.post('%s%s'%(self.base_url, url), params)
            ret = json.loads(response.text)
        except ValueError as e:
            print(e)

        return ret

    def get(self, url):
        "Get"
        ret = None
        try:
            response = self.currentSesion.get('%s%s'%(self.base_url, url))
            ret = json.loads(response.text)
        except ValueError as e:
            print(e)

        return ret

    def put(self, url, params):
        "Put"
        ret = None
        try:
            response = self.currentSesion.put('%s%s'%(self.base_url, url), params)
            ret = json.loads(response.text)
        except ValueError as e:
            print(e)

        return ret

    def delete(self, url):
        "Delete"
        ret = None
        try:
            response = self.currentSesion.delete('%s%s'%(self.base_url, url))
            ret = json.loads(response.text)
        except ValueError as e:
            print(e)

        return ret

    def upload(self, url, params):
        "Upload"
        ret = None
        try:
            response = self.currentSesion.post('%s%s'%(self.base_url, url), params)
            ret = json.loads(response.text)
        except ValueError as e:
            print(e)

        return ret
