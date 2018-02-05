"MagicSession"

import requests
import json
class MagicSession(object):
    "MagicSession"
    def __init__(self, base_url):
        self.currentSesion = requests.Session()
        print("MagicSession construct")
        self.base_url = base_url

    def __del__(self):
        print("MagicSession destruct")

    def post(self, url, params):
        "Post"
        ret = None
        try:
            result = self.currentSesion.post('%s%s'%(self.base_url, url), params)
            print(result.text)
            ret = json.loads(result.text)
        except ValueError as e:
            print(e)

        return ret

    def get(self, url):
        "Get"
        ret = None
        try:
            result = self.currentSesion.get('%s%s'%(self.base_url, url))
            print(result.text)
            ret = json.loads(result.text)
        except ValueError as e:
            print(e)

        return ret

    def put(self, url, params):
        "Put"
        ret = None
        try:
            result = self.currentSesion.put('%s%s'%(self.base_url, url), params)
            print(result.text)
            ret = json.loads(result.text)
        except ValueError as e:
            print(e)

        return ret

    def delete(self, url):
        "Delete"
        ret = None
        try:
            result = self.currentSesion.delete('%s%s'%(self.base_url, url))
            print(result.text)
            ret = json.loads(result.text)
        except ValueError as e:
            print(e)

        return ret

    def upload(self, url, params):
        "Upload"
        ret = None
        try:
            result = self.currentSesion.post('%s%s'%(self.base_url, url), params)
            print(result.text)
            ret = json.loads(result.text)
        except ValueError as e:
            print(e)

        return ret
