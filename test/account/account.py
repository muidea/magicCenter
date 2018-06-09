'account manager'

import random

from account import user
from account import group
from mock import common
from session import session
from cas import cas

class Account:
    'Account Manager'
    def __init__(self, work_session, auth_token):
        self.current_user_session = user.User(work_session, auth_token)
        self.current_group_session = group.Group(work_session, auth_token)
        self.users = {}
        self.groups = {}

        self.__refresh_user()
        self.__refresh_group()

    def __refresh_user(self):
        all_user = {}
        users = self.current_user_session.find_all()
        for _, val in  enumerate(users):
            all_user[val['id']] = val

        self.users = all_user

    def __refresh_group(self):
        all_group = {}
        groups = self.current_group_session.find_all()
        for _, val in enumerate(groups):
            all_group[val['id']] = val

        self.groups = all_group

    def __mock_group(self):
        catalog = common.picker_dict(self.groups, 1)
        catalog = {}
        for _, val in catalog.items():
            catalog = val

        return common.word(), common.paragraph(), catalog

    def __mock_user(self):
        group_num = random.randint(1, 10)
        catalog = common.picker_dict(self.groups, group_num)
        groups = []
        for _, val in catalog.items():
            groups.append(val)

        return common.word(), common.email(), groups

    def verify(self):
        'verify account infomation between local and server'
        local_user = self.users
        self.__refresh_user()
        server_user = self.users

        local_count = len(local_user)
        server_count = len(server_user)

        if local_count > server_count:
            for k in enumerate(local_user.keys()):
                if k not in server_user:
                    print('cant\'t find user in server_user, id:%d', k)
                    break
        elif local_count < server_count:
            for k in enumerate(server_user.keys()):
                if k not in local_user:
                    print('cant\'t find user in local_user, id:%d', k)
                    break

        local_group = self.groups
        self.__refresh_group()
        server_group = self.groups

        local_count = len(local_group)
        server_count = len(server_group)

        if local_count > server_count:
            for k in enumerate(local_group.keys()):
                if k not in server_group:
                    print('cant\'t find user in server_group, id:%d', k)
                    break
        elif local_count < server_count:
            for k in enumerate(server_group.keys()):
                if k not in local_group:
                    print('cant\'t find user in local_group, id:%d', k)
                    break


    def refresh(self, auth_token):
        'refresh account info'
        self.current_user_session.refresh_token(auth_token)
        self.current_group_session.refresh_token(auth_token)

        self.__refresh_user()
        self.__refresh_group()

    def remove(self):
        'remove'
        user_num = random.randint(1, 10)
        user_dict = common.picker_dict(self.users, user_num)
        for key in user_dict.keys():
            if key > 0:
                if self.current_user_session.destroy(key):
                    self.users.pop(key)

        group_num = random.randint(1, 10)
        group_dict = common.picker_dict(self.groups, group_num)
        for key in group_dict.keys():
            if key > 0:
                if self.current_group_session.destroy(key):
                    self.groups.pop(key)

    def clear(self):
        'clear'
        for key in iter(self.users.keys()):
            if key > 0:
                self.current_user_session.destroy(key)

        for key in iter(self.groups.keys()):
            if key > 0:
                self.current_group_session.destroy(key)

    def mock(self):
        'mock'
        group_count = random.randint(1, 10)
        while group_count > 0:
            group_count = group_count - 1
            name, desc, catalog = self.__mock_group()
            new_group = self.current_group_session.create(name, desc, catalog)
            if new_group is None:
                print('create group failed')
                break
            else:
                self.groups[new_group['id']] = new_group

        user_count = random.randint(1, 10)
        while user_count > 0:
            user_count = user_count -1
            account, email, catalog = self.__mock_user()
            new_user = self.current_user_session.create(account, '123', email, catalog)
            if new_user is None:
                print('create user failed')
                break
            else:
                self.users[new_user['id']] = new_user

def main():
    'main'
    work_session = session.MagicSession('http://localhost:8888')
    cas_session = cas.Cas(work_session)
    if not cas_session.login('admin@muidea.com', '123'):
        print('login failed')
    else:
        app = Account(work_session, cas_session.authority_token)
        app.refresh(cas_session.authority_token)
        app.mock()
        app.remove()
        app.verify()
        app.clear()

        cas_session.logout(cas_session.authority_token)
