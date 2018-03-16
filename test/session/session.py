"MagicSession"

import json
import requests
class MagicSession(object):
    "MagicSession"
    def __init__(self, base_url):
        self.current_sesion = requests.Session()
        self.base_url = base_url

    def post(self, url, params):
        "Post"
        ret = None
        try:
            response = self.current_sesion.post('%s%s'%(self.base_url, url), json=params)
            ret = json.loads(response.text)
        except ValueError as except_value:
            print(except_value)

        return ret

    def get(self, url):
        "Get"
        ret = None
        try:
            response = self.current_sesion.get('%s%s'%(self.base_url, url))
            ret = json.loads(response.text)
        except ValueError as except_value:
            print(except_value)

        return ret

    def put(self, url, params):
        "Put"
        ret = None
        try:
            response = self.current_sesion.put('%s%s'%(self.base_url, url), json=params)
            ret = json.loads(response.text)
        except ValueError as except_value:
            print(except_value)

        return ret

    def delete(self, url):
        "Delete"
        ret = None
        try:
            response = self.current_sesion.delete('%s%s'%(self.base_url, url))
            ret = json.loads(response.text)
        except ValueError as except_value:
            print(except_value)

        return ret

    def upload(self, url, params):
        "Upload"
        ret = None
        try:
            response = self.current_sesion.post('%s%s'%(self.base_url, url), params)
            ret = json.loads(response.text)
        except ValueError as except_value:
            print(except_value)

        return ret
